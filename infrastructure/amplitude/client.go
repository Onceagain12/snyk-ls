/*
 * © 2022 Snyk Limited All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package amplitude

import (
	"strings"

	"github.com/amplitude/analytics-go/amplitude"
	"github.com/amplitude/analytics-go/amplitude/plugins/destination"
	"github.com/rs/zerolog/log"

	"github.com/snyk/snyk-ls/ampli"
	"github.com/snyk/snyk-ls/application/config"
	"github.com/snyk/snyk-ls/domain/observability/error_reporting"
	ux2 "github.com/snyk/snyk-ls/domain/observability/ux"
)

type Client struct {
	authenticatedUserId string
	destination         *SegmentPlugin
	errorReporter       error_reporting.ErrorReporter
	authFunc            func() (string, error)
}

type captureEvent func(userId string, eventOptions ...ampli.EventOptions)

func NewAmplitudeClient(authFunc func() (string, error), errorReporter error_reporting.ErrorReporter) ux2.Analytics {
	ampliConfig := amplitude.NewConfig("")

	ampli.Instance.Load(ampli.LoadOptions{
		Client: ampli.LoadClientOptions{
			Configuration: ampliConfig,
		},
	})

	segmentPlugin := NewSegmentPlugin()
	ampli.Instance.Client.Remove(destination.NewAmplitudePlugin().Name()) // remove Amplitude as events destination, as we use Segment for it
	ampli.Instance.Client.Add(segmentPlugin)

	client := &Client{
		destination:   segmentPlugin,
		errorReporter: errorReporter,
		authFunc:      authFunc,
	}

	return client
}

func (c *Client) Initialise() {
	go c.captureInstalledEvent()
}

func (c *Client) Shutdown() error {
	return c.destination.Shutdown()
}

func (c *Client) AnalysisIsReady(properties ux2.AnalysisIsReadyProperties) {
	log.Debug().Str("method", "AnalysisIsReady").Msg("analytics enqueued")
	analysisType := ampli.AnalysisIsReadyAnalysisType(properties.AnalysisType)
	ide := ampli.AnalysisIsReadyIde(getIdeProperty())
	result := ampli.AnalysisIsReadyResult(properties.Result)
	conf := config.CurrentConfig()
	event := ampli.AnalysisIsReady.Builder().
		AnalysisType(analysisType).
		Ide(ide).
		Result(result).
		DurationInSeconds(properties.DurationInSeconds).
		FileCount(properties.FileCount).
		OsPlatform(conf.OsPlatform()).
		OsArch(conf.OsArch()).
		RuntimeName(conf.RuntimeName()).
		RuntimeVersion(conf.RuntimeVersion()).
		Build()

	captureFn := func(authenticatedUserId string, eventOptions ...ampli.EventOptions) {
		ampli.Instance.AnalysisIsReady(authenticatedUserId, event, eventOptions...)
	}
	c.enqueueEvent(captureFn)
}

func (c *Client) AnalysisIsTriggered(properties ux2.AnalysisIsTriggeredProperties) {
	log.Debug().Str("method", "AnalysisIsTriggered").Msg("analytics enqueued")
	analysisTypes := make([]string, 0, len(properties.AnalysisType))
	for _, analysisType := range properties.AnalysisType {
		analysisTypes = append(analysisTypes, string(analysisType))
	}
	conf := config.CurrentConfig()
	ide := ampli.AnalysisIsTriggeredIde(getIdeProperty())
	event := ampli.AnalysisIsTriggered.
		Builder().
		AnalysisType(analysisTypes).
		Ide(ide).
		TriggeredByUser(properties.TriggeredByUser).
		OsPlatform(conf.OsPlatform()).
		OsArch(conf.OsArch()).
		RuntimeName(conf.RuntimeName()).
		RuntimeVersion(conf.RuntimeVersion()).
		Build()

	captureFn := func(authenticatedUserId string, eventOptions ...ampli.EventOptions) {
		ampli.Instance.AnalysisIsTriggered(authenticatedUserId, event, eventOptions...)
	}
	c.enqueueEvent(captureFn)
}

func (c *Client) IssueHoverIsDisplayed(properties ux2.IssueHoverIsDisplayedProperties) {
	log.Debug().Str("method", "IssueHoverIsDisplayed").Msg("analytics enqueued")
	ide := ampli.IssueHoverIsDisplayedIde(getIdeProperty())
	issueType := ampli.IssueHoverIsDisplayedIssueType(properties.IssueType)
	severity := ampli.IssueHoverIsDisplayedSeverity(properties.Severity)
	conf := config.CurrentConfig()
	event := ampli.IssueHoverIsDisplayed.Builder().
		Ide(ide).
		IssueId(properties.IssueId).
		IssueType(issueType).
		Severity(severity).
		OsPlatform(conf.OsPlatform()).
		OsArch(conf.OsArch()).
		RuntimeName(conf.RuntimeName()).
		RuntimeVersion(conf.RuntimeVersion()).
		Build()

	captureFn := func(authenticatedUserId string, eventOptions ...ampli.EventOptions) {
		ampli.Instance.IssueHoverIsDisplayed(authenticatedUserId, event, eventOptions...)
	}
	c.enqueueEvent(captureFn)
}

func (c *Client) PluginIsInstalled(_ ux2.PluginIsInstalledProperties) {
	log.Debug().Str("method", "PluginIsInstalled").Msg("analytics enqueued")
	conf := config.CurrentConfig()

	ide := ampli.PluginIsInstalledIde(getIdeProperty())
	event := ampli.PluginIsInstalled.Builder().
		Ide(ide).
		OsPlatform(conf.OsPlatform()).
		OsArch(conf.OsArch()).
		RuntimeName(conf.RuntimeName()).
		RuntimeVersion(conf.RuntimeVersion()).
		Build()

	captureFn := func(_ string, eventOptions ...ampli.EventOptions) {
		ampli.Instance.PluginIsInstalled("", event, eventOptions...)
	}
	c.enqueueEvent(captureFn)
}

func (c *Client) ScanModeIsSelected(properties ux2.ScanModeIsSelectedProperties) {
	conf := config.CurrentConfig()
	ide := ampli.ScanModeIsSelectedIde(getIdeProperty())

	event := ampli.ScanModeIsSelected.Builder().
		Ide(ide).
		ScanMode(ampli.ScanModeIsSelectedScanMode(properties.ScanningMode)).
		OsPlatform(conf.OsPlatform()).
		OsArch(conf.OsArch()).
		RuntimeName(conf.RuntimeName()).
		RuntimeVersion(conf.RuntimeVersion()).
		Build()

	captureFn := func(authenticatedUserId string, eventOptions ...ampli.EventOptions) {
		ampli.Instance.ScanModeIsSelected(authenticatedUserId, event, eventOptions...)
	}
	c.enqueueEvent(captureFn)
}

func (c *Client) enqueueEvent(eventFn captureEvent) {
	if config.CurrentConfig().IsTelemetryEnabled() {
		eventFn(
			c.authenticatedUserId,
			ampli.EventOptions{
				DeviceID: config.CurrentConfig().DeviceID(),
			})
	}
}

func (c *Client) Identify() {
	method := "infrastructure.segment.client"
	log.Debug().Str("method", method).Msg("Identifying a user.")
	if !config.CurrentConfig().NonEmptyToken() {
		c.authenticatedUserId = ""
		return
	}

	userId, err := c.authFunc()
	if err != nil {
		log.Debug().Str("method", method).Err(err).Msg("Failed to identify user.")
		return
	}
	c.authenticatedUserId = userId

	if !config.CurrentConfig().IsTelemetryEnabled() {
		return
	}

	identifyEvent := ampli.Identify.Builder().UserId(userId).Build()
	ampli.Instance.Identify(c.authenticatedUserId, identifyEvent, ampli.EventOptions{})
}

// Only return an IDE property if it's a recognized IDE in the tracking plan
func getIdeProperty() ux2.IDE {
	// Standardize the names
	integrationName := strings.Replace(strings.ToLower(config.CurrentConfig().IntegrationName()), "_", " ", -1)

	switch integrationName {
	case string(ux2.Eclipse):
		return ux2.Eclipse
	case string("vs code"):
		return ux2.VisualStudioCode
	case string(ux2.VisualStudio):
		return ux2.VisualStudio
	case string(ux2.JetBrains):
		return ux2.JetBrains
	default:
		return "Other" // ensure we pass Amplitude validation, when IDE is not in the tracking plan
	}
}

type segmentLogger struct{}

func (s *segmentLogger) Logf(format string, args ...any) {
	log.Debug().Msgf(format, args...)
}

func (s *segmentLogger) Errorf(format string, args ...any) {
	log.Error().Msgf(format, args...)
}
