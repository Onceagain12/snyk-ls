// ampli.go
//
// Ampli - A strong typed wrapper for your Analytics
//
// This file is generated by Amplitude.
// To update run 'ampli pull language-server'
//
// Required dependencies: github.com/amplitude/analytics-go/amplitude@v0.0.5
// Tracking Plan Version: 247
// Build: 1.0.0
// Runtime: go-ampli
//
// View Tracking Plan: https://data.amplitude.com/snyk/Snyk/events/main/latest
//
// Full Setup Instructions: https://data.amplitude.com/snyk/Snyk/implementation/main/latest/getting-started/language-server
//

package ampli

import (
	"log"
	"sync"

	"github.com/amplitude/analytics-go/amplitude"
)

type (
	EventOptions  = amplitude.EventOptions
	ExecuteResult = amplitude.ExecuteResult
)

const (
	IdentifyEventType      = amplitude.IdentifyEventType
	GroupIdentifyEventType = amplitude.GroupIdentifyEventType

	ServerZoneUS = amplitude.ServerZoneUS
	ServerZoneEU = amplitude.ServerZoneEU
)

var (
	NewClientConfig = amplitude.NewConfig
	NewClient       = amplitude.NewClient
)

var Instance = Ampli{}

type Environment string

const (
	EnvironmentProduction Environment = `production`

	EnvironmentDevelopment Environment = `development`
)

var APIKey = map[Environment]string{
	EnvironmentProduction: `af27ef9ddacb424e34b42983161c1017`,

	EnvironmentDevelopment: `0685e531d804b43d914d4857cbaeab49`,
}

// LoadClientOptions is Client options setting to initialize Ampli client.
//
// Params:
//   - APIKey: the API key of Amplitude project
//   - Instance: the core SDK instance used by Ampli client
//   - Configuration: the core SDK client configuration instance
type LoadClientOptions struct {
	APIKey        string
	Instance      amplitude.Client
	Configuration amplitude.Config
}

// LoadOptions is options setting to initialize Ampli client.
//
// Params:
//   - Environment: the environment of Amplitude Data project
//   - Disabled: the flag of disabled Ampli client
//   - Client: the LoadClientOptions struct
type LoadOptions struct {
	Environment Environment
	Disabled    bool
	Client      LoadClientOptions
}

type baseEvent struct {
	eventType  string
	properties map[string]interface{}
}

type Event interface {
	ToAmplitudeEvent() amplitude.Event
}

func newBaseEvent(eventType string, properties map[string]interface{}) baseEvent {
	return baseEvent{
		eventType:  eventType,
		properties: properties,
	}
}

func (event baseEvent) ToAmplitudeEvent() amplitude.Event {
	return amplitude.Event{
		EventType:       event.eventType,
		EventProperties: event.properties,
	}
}

type IdentifyAccountType string

var Identify = struct {
	AccountType struct {
		User IdentifyAccountType

		Service IdentifyAccountType

		AppInstance IdentifyAccountType

		AutomatedTestUser IdentifyAccountType
	}
	Builder func() IdentifyBuilder
}{
	AccountType: struct {
		User IdentifyAccountType

		Service IdentifyAccountType

		AppInstance IdentifyAccountType

		AutomatedTestUser IdentifyAccountType
	}{
		User: `user`,

		Service: `service`,

		AppInstance: `app-instance`,

		AutomatedTestUser: `automated-test-user`,
	},
	Builder: func() IdentifyBuilder {
		return &identifyBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type IdentifyEvent interface {
	Event
	identify()
}

type identifyEvent struct {
	baseEvent
}

func (e identifyEvent) identify() {
}

type IdentifyBuilder interface {
	Build() IdentifyEvent
	AccountType(accountType IdentifyAccountType) IdentifyBuilder
	AdminLink(adminLink string) IdentifyBuilder
	AuthProvider(authProvider string) IdentifyBuilder
	CreatedAt(createdAt interface{}) IdentifyBuilder
	Email(email string) IdentifyBuilder
	HasFirstIntegration(hasFirstIntegration bool) IdentifyBuilder
	HasFirstProject(hasFirstProject bool) IdentifyBuilder
	HasPersonalEmail(hasPersonalEmail bool) IdentifyBuilder
	IsAppUiBetaEnabled(isAppUiBetaEnabled bool) IdentifyBuilder
	IsNonUser(isNonUser bool) IdentifyBuilder
	IsSnyk(isSnyk bool) IdentifyBuilder
	IsSnykAdmin(isSnykAdmin bool) IdentifyBuilder
	LearnPreferredEcosystems(learnPreferredEcosystems []string) IdentifyBuilder
	ProductUpdatesConsent(productUpdatesConsent bool) IdentifyBuilder
	ProductUpdatesConsentIsDisplayed(productUpdatesConsentIsDisplayed bool) IdentifyBuilder
	UserId(userId string) IdentifyBuilder
	Username(username string) IdentifyBuilder
	UtmCampaign(utmCampaign string) IdentifyBuilder
	UtmContent(utmContent string) IdentifyBuilder
	UtmMedium(utmMedium string) IdentifyBuilder
	UtmSource(utmSource string) IdentifyBuilder
	UtmTerm(utmTerm string) IdentifyBuilder
}

type identifyBuilder struct {
	properties map[string]interface{}
}

func (b *identifyBuilder) AccountType(accountType IdentifyAccountType) IdentifyBuilder {
	b.properties[`accountType`] = accountType

	return b
}

func (b *identifyBuilder) AdminLink(adminLink string) IdentifyBuilder {
	b.properties[`adminLink`] = adminLink

	return b
}

func (b *identifyBuilder) AuthProvider(authProvider string) IdentifyBuilder {
	b.properties[`authProvider`] = authProvider

	return b
}

func (b *identifyBuilder) CreatedAt(createdAt interface{}) IdentifyBuilder {
	b.properties[`createdAt`] = createdAt

	return b
}

func (b *identifyBuilder) Email(email string) IdentifyBuilder {
	b.properties[`email`] = email

	return b
}

func (b *identifyBuilder) HasFirstIntegration(hasFirstIntegration bool) IdentifyBuilder {
	b.properties[`hasFirstIntegration`] = hasFirstIntegration

	return b
}

func (b *identifyBuilder) HasFirstProject(hasFirstProject bool) IdentifyBuilder {
	b.properties[`hasFirstProject`] = hasFirstProject

	return b
}

func (b *identifyBuilder) HasPersonalEmail(hasPersonalEmail bool) IdentifyBuilder {
	b.properties[`hasPersonalEmail`] = hasPersonalEmail

	return b
}

func (b *identifyBuilder) IsAppUiBetaEnabled(isAppUiBetaEnabled bool) IdentifyBuilder {
	b.properties[`isAppUIBetaEnabled`] = isAppUiBetaEnabled

	return b
}

func (b *identifyBuilder) IsNonUser(isNonUser bool) IdentifyBuilder {
	b.properties[`isNonUser`] = isNonUser

	return b
}

func (b *identifyBuilder) IsSnyk(isSnyk bool) IdentifyBuilder {
	b.properties[`isSnyk`] = isSnyk

	return b
}

func (b *identifyBuilder) IsSnykAdmin(isSnykAdmin bool) IdentifyBuilder {
	b.properties[`isSnykAdmin`] = isSnykAdmin

	return b
}

func (b *identifyBuilder) LearnPreferredEcosystems(learnPreferredEcosystems []string) IdentifyBuilder {
	b.properties[`learnPreferredEcosystems`] = learnPreferredEcosystems

	return b
}

func (b *identifyBuilder) ProductUpdatesConsent(productUpdatesConsent bool) IdentifyBuilder {
	b.properties[`productUpdatesConsent`] = productUpdatesConsent

	return b
}

func (b *identifyBuilder) ProductUpdatesConsentIsDisplayed(productUpdatesConsentIsDisplayed bool) IdentifyBuilder {
	b.properties[`productUpdatesConsentIsDisplayed`] = productUpdatesConsentIsDisplayed

	return b
}

func (b *identifyBuilder) UserId(userId string) IdentifyBuilder {
	b.properties[`user_id`] = userId

	return b
}

func (b *identifyBuilder) Username(username string) IdentifyBuilder {
	b.properties[`username`] = username

	return b
}

func (b *identifyBuilder) UtmCampaign(utmCampaign string) IdentifyBuilder {
	b.properties[`utm_campaign`] = utmCampaign

	return b
}

func (b *identifyBuilder) UtmContent(utmContent string) IdentifyBuilder {
	b.properties[`utm_content`] = utmContent

	return b
}

func (b *identifyBuilder) UtmMedium(utmMedium string) IdentifyBuilder {
	b.properties[`utm_medium`] = utmMedium

	return b
}

func (b *identifyBuilder) UtmSource(utmSource string) IdentifyBuilder {
	b.properties[`utm_source`] = utmSource

	return b
}

func (b *identifyBuilder) UtmTerm(utmTerm string) IdentifyBuilder {
	b.properties[`utm_term`] = utmTerm

	return b
}

func (b *identifyBuilder) Build() IdentifyEvent {
	return &identifyEvent{
		newBaseEvent(`Identify`, b.properties),
	}
}

func (e identifyEvent) ToAmplitudeEvent() amplitude.Event {
	identify := amplitude.Identify{}
	for name, value := range e.properties {
		identify.Set(name, value)
	}

	return amplitude.Event{
		EventType:      IdentifyEventType,
		UserProperties: identify.Properties,
	}
}

type GroupGroupType string

var Group = struct {
	GroupType struct {
		Org GroupGroupType

		Group GroupGroupType

		Account GroupGroupType
	}
	Builder func() GroupBuilder
}{
	GroupType: struct {
		Org GroupGroupType

		Group GroupGroupType

		Account GroupGroupType
	}{
		Org: `org`,

		Group: `group`,

		Account: `account`,
	},
	Builder: func() GroupBuilder {
		return &groupBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type GroupEvent interface {
	Event
	group()
}

type groupEvent struct {
	baseEvent
}

func (e groupEvent) group() {
}

type GroupBuilder interface {
	Build() GroupEvent
	AmplitudeGroupId(amplitudeGroupId interface{}) GroupBuilder
	AmplitudeGroupName(amplitudeGroupName interface{}) GroupBuilder
	Set(set interface{}) GroupBuilder
	Unset(unset interface{}) GroupBuilder
	AccountArr(accountArr interface{}) GroupBuilder
	AccountPlan(accountPlan interface{}) GroupBuilder
	BillingFrequency(billingFrequency interface{}) GroupBuilder
	CodeLicenses(codeLicenses interface{}) GroupBuilder
	ContainerLicenses(containerLicenses interface{}) GroupBuilder
	CountFixesFirst30Days(countFixesFirst30Days interface{}) GroupBuilder
	CountFixesFirst7Days(countFixesFirst7Days interface{}) GroupBuilder
	CountFixesPast30Days(countFixesPast30Days interface{}) GroupBuilder
	CountFixesPast7Days(countFixesPast7Days interface{}) GroupBuilder
	CountFixesTotal(countFixesTotal interface{}) GroupBuilder
	CurrentEngagementState(currentEngagementState interface{}) GroupBuilder
	DateLastEngagementStateChange(dateLastEngagementStateChange interface{}) GroupBuilder
	DaysSinceLastEngagementStateChange(daysSinceLastEngagementStateChange interface{}) GroupBuilder
	DbFeedLicenses(dbFeedLicenses interface{}) GroupBuilder
	FreeTrialEndDate(freeTrialEndDate interface{}) GroupBuilder
	FreeTrialStartDate(freeTrialStartDate interface{}) GroupBuilder
	GroupId(groupId string) GroupBuilder
	GroupName(groupName string) GroupBuilder
	GroupType(groupType GroupGroupType) GroupBuilder
	HasFixFirst30Days(hasFixFirst30Days interface{}) GroupBuilder
	HasFixFirst7Days(hasFixFirst7Days interface{}) GroupBuilder
	HasFixPast30Days(hasFixPast30Days interface{}) GroupBuilder
	HasFixPast7Days(hasFixPast7Days interface{}) GroupBuilder
	IacLicenses(iacLicenses interface{}) GroupBuilder
	Id(id interface{}) GroupBuilder
	InternalName(internalName interface{}) GroupBuilder
	IsPassthrough(isPassthrough interface{}) GroupBuilder
	Name(name interface{}) GroupBuilder
	OpenSourceLicenses(openSourceLicenses interface{}) GroupBuilder
	Plan(plan string) GroupBuilder
	PriorEngagementState(priorEngagementState interface{}) GroupBuilder
	ProjectTypes(projectTypes []string) GroupBuilder
}

type groupBuilder struct {
	properties map[string]interface{}
}

func (b *groupBuilder) AmplitudeGroupId(amplitudeGroupId interface{}) GroupBuilder {
	b.properties[`[Amplitude] Group ID`] = amplitudeGroupId

	return b
}

func (b *groupBuilder) AmplitudeGroupName(amplitudeGroupName interface{}) GroupBuilder {
	b.properties[`[Amplitude] Group name`] = amplitudeGroupName

	return b
}

func (b *groupBuilder) Set(set interface{}) GroupBuilder {
	b.properties[`$set`] = set

	return b
}

func (b *groupBuilder) Unset(unset interface{}) GroupBuilder {
	b.properties[`$unset`] = unset

	return b
}

func (b *groupBuilder) AccountArr(accountArr interface{}) GroupBuilder {
	b.properties[`Account ARR`] = accountArr

	return b
}

func (b *groupBuilder) AccountPlan(accountPlan interface{}) GroupBuilder {
	b.properties[`Account Plan`] = accountPlan

	return b
}

func (b *groupBuilder) BillingFrequency(billingFrequency interface{}) GroupBuilder {
	b.properties[`Billing Frequency`] = billingFrequency

	return b
}

func (b *groupBuilder) CodeLicenses(codeLicenses interface{}) GroupBuilder {
	b.properties[`Code Licenses`] = codeLicenses

	return b
}

func (b *groupBuilder) ContainerLicenses(containerLicenses interface{}) GroupBuilder {
	b.properties[`Container Licenses`] = containerLicenses

	return b
}

func (b *groupBuilder) CountFixesFirst30Days(countFixesFirst30Days interface{}) GroupBuilder {
	b.properties[`countFixesFirst30Days`] = countFixesFirst30Days

	return b
}

func (b *groupBuilder) CountFixesFirst7Days(countFixesFirst7Days interface{}) GroupBuilder {
	b.properties[`countFixesFirst7Days`] = countFixesFirst7Days

	return b
}

func (b *groupBuilder) CountFixesPast30Days(countFixesPast30Days interface{}) GroupBuilder {
	b.properties[`countFixesPast30Days`] = countFixesPast30Days

	return b
}

func (b *groupBuilder) CountFixesPast7Days(countFixesPast7Days interface{}) GroupBuilder {
	b.properties[`countFixesPast7Days`] = countFixesPast7Days

	return b
}

func (b *groupBuilder) CountFixesTotal(countFixesTotal interface{}) GroupBuilder {
	b.properties[`countFixesTotal`] = countFixesTotal

	return b
}

func (b *groupBuilder) CurrentEngagementState(currentEngagementState interface{}) GroupBuilder {
	b.properties[`currentEngagementState`] = currentEngagementState

	return b
}

func (b *groupBuilder) DateLastEngagementStateChange(dateLastEngagementStateChange interface{}) GroupBuilder {
	b.properties[`dateLastEngagementStateChange`] = dateLastEngagementStateChange

	return b
}

func (b *groupBuilder) DaysSinceLastEngagementStateChange(daysSinceLastEngagementStateChange interface{}) GroupBuilder {
	b.properties[`daysSinceLastEngagementStateChange`] = daysSinceLastEngagementStateChange

	return b
}

func (b *groupBuilder) DbFeedLicenses(dbFeedLicenses interface{}) GroupBuilder {
	b.properties[`DB Feed Licenses`] = dbFeedLicenses

	return b
}

func (b *groupBuilder) FreeTrialEndDate(freeTrialEndDate interface{}) GroupBuilder {
	b.properties[`Free Trial End Date`] = freeTrialEndDate

	return b
}

func (b *groupBuilder) FreeTrialStartDate(freeTrialStartDate interface{}) GroupBuilder {
	b.properties[`Free Trial Start Date`] = freeTrialStartDate

	return b
}

func (b *groupBuilder) GroupId(groupId string) GroupBuilder {
	b.properties[`groupId`] = groupId

	return b
}

func (b *groupBuilder) GroupName(groupName string) GroupBuilder {
	b.properties[`groupName`] = groupName

	return b
}

func (b *groupBuilder) GroupType(groupType GroupGroupType) GroupBuilder {
	b.properties[`groupType`] = groupType

	return b
}

func (b *groupBuilder) HasFixFirst30Days(hasFixFirst30Days interface{}) GroupBuilder {
	b.properties[`hasFixFirst30Days`] = hasFixFirst30Days

	return b
}

func (b *groupBuilder) HasFixFirst7Days(hasFixFirst7Days interface{}) GroupBuilder {
	b.properties[`hasFixFirst7Days`] = hasFixFirst7Days

	return b
}

func (b *groupBuilder) HasFixPast30Days(hasFixPast30Days interface{}) GroupBuilder {
	b.properties[`hasFixPast30Days`] = hasFixPast30Days

	return b
}

func (b *groupBuilder) HasFixPast7Days(hasFixPast7Days interface{}) GroupBuilder {
	b.properties[`hasFixPast7Days`] = hasFixPast7Days

	return b
}

func (b *groupBuilder) IacLicenses(iacLicenses interface{}) GroupBuilder {
	b.properties[`IAC Licenses`] = iacLicenses

	return b
}

func (b *groupBuilder) Id(id interface{}) GroupBuilder {
	b.properties[`id`] = id

	return b
}

func (b *groupBuilder) InternalName(internalName interface{}) GroupBuilder {
	b.properties[`internalName`] = internalName

	return b
}

func (b *groupBuilder) IsPassthrough(isPassthrough interface{}) GroupBuilder {
	b.properties[`isPassthrough`] = isPassthrough

	return b
}

func (b *groupBuilder) Name(name interface{}) GroupBuilder {
	b.properties[`name`] = name

	return b
}

func (b *groupBuilder) OpenSourceLicenses(openSourceLicenses interface{}) GroupBuilder {
	b.properties[`Open Source Licenses`] = openSourceLicenses

	return b
}

func (b *groupBuilder) Plan(plan string) GroupBuilder {
	b.properties[`plan`] = plan

	return b
}

func (b *groupBuilder) PriorEngagementState(priorEngagementState interface{}) GroupBuilder {
	b.properties[`priorEngagementState`] = priorEngagementState

	return b
}

func (b *groupBuilder) ProjectTypes(projectTypes []string) GroupBuilder {
	b.properties[`projectTypes`] = projectTypes

	return b
}

func (b *groupBuilder) Build() GroupEvent {
	return &groupEvent{
		newBaseEvent(`Group`, b.properties),
	}
}

func (e groupEvent) ToAmplitudeEvent() amplitude.Event {
	identify := amplitude.Identify{}
	for name, value := range e.properties {
		identify.Set(name, value)
	}

	return amplitude.Event{
		EventType:       GroupIdentifyEventType,
		GroupProperties: identify.Properties,
	}
}

type AnalysisIsReadyAnalysisType string

type AnalysisIsReadyIde string

type AnalysisIsReadyResult string

var AnalysisIsReady = struct {
	AnalysisType struct {
		SnykAdvisor AnalysisIsReadyAnalysisType

		SnykCodeQuality AnalysisIsReadyAnalysisType

		SnykCodeSecurity AnalysisIsReadyAnalysisType

		SnykOpenSource AnalysisIsReadyAnalysisType

		SnykContainer AnalysisIsReadyAnalysisType

		SnykInfrastructureAsCode AnalysisIsReadyAnalysisType
	}
	Ide struct {
		VisualStudioCode AnalysisIsReadyIde

		VisualStudio AnalysisIsReadyIde

		Eclipse AnalysisIsReadyIde

		JetBrains AnalysisIsReadyIde

		Other AnalysisIsReadyIde
	}
	Result struct {
		Success AnalysisIsReadyResult

		Error AnalysisIsReadyResult
	}
	Builder func() interface {
		AnalysisType(analysisType AnalysisIsReadyAnalysisType) interface {
			Ide(ide AnalysisIsReadyIde) interface {
				Result(result AnalysisIsReadyResult) AnalysisIsReadyBuilder
			}
		}
	}
}{
	AnalysisType: struct {
		SnykAdvisor AnalysisIsReadyAnalysisType

		SnykCodeQuality AnalysisIsReadyAnalysisType

		SnykCodeSecurity AnalysisIsReadyAnalysisType

		SnykOpenSource AnalysisIsReadyAnalysisType

		SnykContainer AnalysisIsReadyAnalysisType

		SnykInfrastructureAsCode AnalysisIsReadyAnalysisType
	}{
		SnykAdvisor: `Snyk Advisor`,

		SnykCodeQuality: `Snyk Code Quality`,

		SnykCodeSecurity: `Snyk Code Security`,

		SnykOpenSource: `Snyk Open Source`,

		SnykContainer: `Snyk Container`,

		SnykInfrastructureAsCode: `Snyk Infrastructure as Code`,
	},
	Ide: struct {
		VisualStudioCode AnalysisIsReadyIde

		VisualStudio AnalysisIsReadyIde

		Eclipse AnalysisIsReadyIde

		JetBrains AnalysisIsReadyIde

		Other AnalysisIsReadyIde
	}{
		VisualStudioCode: `Visual Studio Code`,

		VisualStudio: `Visual Studio`,

		Eclipse: `Eclipse`,

		JetBrains: `JetBrains`,

		Other: `Other`,
	},
	Result: struct {
		Success AnalysisIsReadyResult

		Error AnalysisIsReadyResult
	}{
		Success: `Success`,

		Error: `Error`,
	},
	Builder: func() interface {
		AnalysisType(analysisType AnalysisIsReadyAnalysisType) interface {
			Ide(ide AnalysisIsReadyIde) interface {
				Result(result AnalysisIsReadyResult) AnalysisIsReadyBuilder
			}
		}
	} {
		return &analysisIsReadyBuilder{
			properties: map[string]interface{}{
				`itly`: true,
			},
		}
	},
}

type AnalysisIsReadyEvent interface {
	Event
	analysisIsReady()
}

type analysisIsReadyEvent struct {
	baseEvent
}

func (e analysisIsReadyEvent) analysisIsReady() {
}

type AnalysisIsReadyBuilder interface {
	Build() AnalysisIsReadyEvent
	DurationInSeconds(durationInSeconds float64) AnalysisIsReadyBuilder
	FileCount(fileCount int) AnalysisIsReadyBuilder
	OsArch(osArch string) AnalysisIsReadyBuilder
	OsPlatform(osPlatform string) AnalysisIsReadyBuilder
	RuntimeName(runtimeName string) AnalysisIsReadyBuilder
	RuntimeVersion(runtimeVersion string) AnalysisIsReadyBuilder
}

type analysisIsReadyBuilder struct {
	properties map[string]interface{}
}

func (b *analysisIsReadyBuilder) AnalysisType(analysisType AnalysisIsReadyAnalysisType) interface {
	Ide(ide AnalysisIsReadyIde) interface {
		Result(result AnalysisIsReadyResult) AnalysisIsReadyBuilder
	}
} {
	b.properties[`analysisType`] = analysisType

	return b
}

func (b *analysisIsReadyBuilder) Ide(ide AnalysisIsReadyIde) interface {
	Result(result AnalysisIsReadyResult) AnalysisIsReadyBuilder
} {
	b.properties[`ide`] = ide

	return b
}

func (b *analysisIsReadyBuilder) Result(result AnalysisIsReadyResult) AnalysisIsReadyBuilder {
	b.properties[`result`] = result

	return b
}

func (b *analysisIsReadyBuilder) DurationInSeconds(durationInSeconds float64) AnalysisIsReadyBuilder {
	b.properties[`durationInSeconds`] = durationInSeconds

	return b
}

func (b *analysisIsReadyBuilder) FileCount(fileCount int) AnalysisIsReadyBuilder {
	b.properties[`fileCount`] = fileCount

	return b
}

func (b *analysisIsReadyBuilder) OsArch(osArch string) AnalysisIsReadyBuilder {
	b.properties[`osArch`] = osArch

	return b
}

func (b *analysisIsReadyBuilder) OsPlatform(osPlatform string) AnalysisIsReadyBuilder {
	b.properties[`osPlatform`] = osPlatform

	return b
}

func (b *analysisIsReadyBuilder) RuntimeName(runtimeName string) AnalysisIsReadyBuilder {
	b.properties[`runtimeName`] = runtimeName

	return b
}

func (b *analysisIsReadyBuilder) RuntimeVersion(runtimeVersion string) AnalysisIsReadyBuilder {
	b.properties[`runtimeVersion`] = runtimeVersion

	return b
}

func (b *analysisIsReadyBuilder) Build() AnalysisIsReadyEvent {
	return &analysisIsReadyEvent{
		newBaseEvent(`Analysis Is Ready`, b.properties),
	}
}

type AnalysisIsTriggeredIde string

var AnalysisIsTriggered = struct {
	Ide struct {
		VisualStudioCode AnalysisIsTriggeredIde

		VisualStudio AnalysisIsTriggeredIde

		Eclipse AnalysisIsTriggeredIde

		JetBrains AnalysisIsTriggeredIde

		Other AnalysisIsTriggeredIde
	}
	Builder func() interface {
		AnalysisType(analysisType []string) interface {
			Ide(ide AnalysisIsTriggeredIde) interface {
				TriggeredByUser(triggeredByUser bool) AnalysisIsTriggeredBuilder
			}
		}
	}
}{
	Ide: struct {
		VisualStudioCode AnalysisIsTriggeredIde

		VisualStudio AnalysisIsTriggeredIde

		Eclipse AnalysisIsTriggeredIde

		JetBrains AnalysisIsTriggeredIde

		Other AnalysisIsTriggeredIde
	}{
		VisualStudioCode: `Visual Studio Code`,

		VisualStudio: `Visual Studio`,

		Eclipse: `Eclipse`,

		JetBrains: `JetBrains`,

		Other: `Other`,
	},
	Builder: func() interface {
		AnalysisType(analysisType []string) interface {
			Ide(ide AnalysisIsTriggeredIde) interface {
				TriggeredByUser(triggeredByUser bool) AnalysisIsTriggeredBuilder
			}
		}
	} {
		return &analysisIsTriggeredBuilder{
			properties: map[string]interface{}{
				`itly`: true,
			},
		}
	},
}

type AnalysisIsTriggeredEvent interface {
	Event
	analysisIsTriggered()
}

type analysisIsTriggeredEvent struct {
	baseEvent
}

func (e analysisIsTriggeredEvent) analysisIsTriggered() {
}

type AnalysisIsTriggeredBuilder interface {
	Build() AnalysisIsTriggeredEvent
	OsArch(osArch string) AnalysisIsTriggeredBuilder
	OsPlatform(osPlatform string) AnalysisIsTriggeredBuilder
	RuntimeName(runtimeName string) AnalysisIsTriggeredBuilder
	RuntimeVersion(runtimeVersion string) AnalysisIsTriggeredBuilder
}

type analysisIsTriggeredBuilder struct {
	properties map[string]interface{}
}

func (b *analysisIsTriggeredBuilder) AnalysisType(analysisType []string) interface {
	Ide(ide AnalysisIsTriggeredIde) interface {
		TriggeredByUser(triggeredByUser bool) AnalysisIsTriggeredBuilder
	}
} {
	b.properties[`analysisType`] = analysisType

	return b
}

func (b *analysisIsTriggeredBuilder) Ide(ide AnalysisIsTriggeredIde) interface {
	TriggeredByUser(triggeredByUser bool) AnalysisIsTriggeredBuilder
} {
	b.properties[`ide`] = ide

	return b
}

func (b *analysisIsTriggeredBuilder) TriggeredByUser(triggeredByUser bool) AnalysisIsTriggeredBuilder {
	b.properties[`triggeredByUser`] = triggeredByUser

	return b
}

func (b *analysisIsTriggeredBuilder) OsArch(osArch string) AnalysisIsTriggeredBuilder {
	b.properties[`osArch`] = osArch

	return b
}

func (b *analysisIsTriggeredBuilder) OsPlatform(osPlatform string) AnalysisIsTriggeredBuilder {
	b.properties[`osPlatform`] = osPlatform

	return b
}

func (b *analysisIsTriggeredBuilder) RuntimeName(runtimeName string) AnalysisIsTriggeredBuilder {
	b.properties[`runtimeName`] = runtimeName

	return b
}

func (b *analysisIsTriggeredBuilder) RuntimeVersion(runtimeVersion string) AnalysisIsTriggeredBuilder {
	b.properties[`runtimeVersion`] = runtimeVersion

	return b
}

func (b *analysisIsTriggeredBuilder) Build() AnalysisIsTriggeredEvent {
	return &analysisIsTriggeredEvent{
		newBaseEvent(`Analysis Is Triggered`, b.properties),
	}
}

type IssueHoverIsDisplayedIde string

type IssueHoverIsDisplayedIssueType string

type IssueHoverIsDisplayedSeverity string

var IssueHoverIsDisplayed = struct {
	Ide struct {
		VisualStudioCode IssueHoverIsDisplayedIde

		VisualStudio IssueHoverIsDisplayedIde

		Eclipse IssueHoverIsDisplayedIde

		JetBrains IssueHoverIsDisplayedIde

		Other IssueHoverIsDisplayedIde
	}
	IssueType struct {
		Advisor IssueHoverIsDisplayedIssueType

		CodeQualityIssue IssueHoverIsDisplayedIssueType

		CodeSecurityVulnerability IssueHoverIsDisplayedIssueType

		LicenceIssue IssueHoverIsDisplayedIssueType

		OpenSourceVulnerability IssueHoverIsDisplayedIssueType

		InfrastructureAsCodeIssue IssueHoverIsDisplayedIssueType

		ContainerVulnerability IssueHoverIsDisplayedIssueType
	}
	Severity struct {
		High IssueHoverIsDisplayedSeverity

		Medium IssueHoverIsDisplayedSeverity

		Low IssueHoverIsDisplayedSeverity

		Critical IssueHoverIsDisplayedSeverity
	}
	Builder func() interface {
		Ide(ide IssueHoverIsDisplayedIde) IssueHoverIsDisplayedBuilder
	}
}{
	Ide: struct {
		VisualStudioCode IssueHoverIsDisplayedIde

		VisualStudio IssueHoverIsDisplayedIde

		Eclipse IssueHoverIsDisplayedIde

		JetBrains IssueHoverIsDisplayedIde

		Other IssueHoverIsDisplayedIde
	}{
		VisualStudioCode: `Visual Studio Code`,

		VisualStudio: `Visual Studio`,

		Eclipse: `Eclipse`,

		JetBrains: `JetBrains`,

		Other: `Other`,
	},
	IssueType: struct {
		Advisor IssueHoverIsDisplayedIssueType

		CodeQualityIssue IssueHoverIsDisplayedIssueType

		CodeSecurityVulnerability IssueHoverIsDisplayedIssueType

		LicenceIssue IssueHoverIsDisplayedIssueType

		OpenSourceVulnerability IssueHoverIsDisplayedIssueType

		InfrastructureAsCodeIssue IssueHoverIsDisplayedIssueType

		ContainerVulnerability IssueHoverIsDisplayedIssueType
	}{
		Advisor: `Advisor`,

		CodeQualityIssue: `Code Quality Issue`,

		CodeSecurityVulnerability: `Code Security Vulnerability`,

		LicenceIssue: `Licence Issue`,

		OpenSourceVulnerability: `Open Source Vulnerability`,

		InfrastructureAsCodeIssue: `Infrastructure as Code Issue`,

		ContainerVulnerability: `Container Vulnerability`,
	},
	Severity: struct {
		High IssueHoverIsDisplayedSeverity

		Medium IssueHoverIsDisplayedSeverity

		Low IssueHoverIsDisplayedSeverity

		Critical IssueHoverIsDisplayedSeverity
	}{
		High: `High`,

		Medium: `Medium`,

		Low: `Low`,

		Critical: `Critical`,
	},
	Builder: func() interface {
		Ide(ide IssueHoverIsDisplayedIde) IssueHoverIsDisplayedBuilder
	} {
		return &issueHoverIsDisplayedBuilder{
			properties: map[string]interface{}{
				`itly`: true,
			},
		}
	},
}

type IssueHoverIsDisplayedEvent interface {
	Event
	issueHoverIsDisplayed()
}

type issueHoverIsDisplayedEvent struct {
	baseEvent
}

func (e issueHoverIsDisplayedEvent) issueHoverIsDisplayed() {
}

type IssueHoverIsDisplayedBuilder interface {
	Build() IssueHoverIsDisplayedEvent
	IssueId(issueId string) IssueHoverIsDisplayedBuilder
	IssueType(issueType IssueHoverIsDisplayedIssueType) IssueHoverIsDisplayedBuilder
	OsArch(osArch string) IssueHoverIsDisplayedBuilder
	OsPlatform(osPlatform string) IssueHoverIsDisplayedBuilder
	RuntimeName(runtimeName string) IssueHoverIsDisplayedBuilder
	RuntimeVersion(runtimeVersion string) IssueHoverIsDisplayedBuilder
	Severity(severity IssueHoverIsDisplayedSeverity) IssueHoverIsDisplayedBuilder
}

type issueHoverIsDisplayedBuilder struct {
	properties map[string]interface{}
}

func (b *issueHoverIsDisplayedBuilder) Ide(ide IssueHoverIsDisplayedIde) IssueHoverIsDisplayedBuilder {
	b.properties[`ide`] = ide

	return b
}

func (b *issueHoverIsDisplayedBuilder) IssueId(issueId string) IssueHoverIsDisplayedBuilder {
	b.properties[`issueId`] = issueId

	return b
}

func (b *issueHoverIsDisplayedBuilder) IssueType(issueType IssueHoverIsDisplayedIssueType) IssueHoverIsDisplayedBuilder {
	b.properties[`issueType`] = issueType

	return b
}

func (b *issueHoverIsDisplayedBuilder) OsArch(osArch string) IssueHoverIsDisplayedBuilder {
	b.properties[`osArch`] = osArch

	return b
}

func (b *issueHoverIsDisplayedBuilder) OsPlatform(osPlatform string) IssueHoverIsDisplayedBuilder {
	b.properties[`osPlatform`] = osPlatform

	return b
}

func (b *issueHoverIsDisplayedBuilder) RuntimeName(runtimeName string) IssueHoverIsDisplayedBuilder {
	b.properties[`runtimeName`] = runtimeName

	return b
}

func (b *issueHoverIsDisplayedBuilder) RuntimeVersion(runtimeVersion string) IssueHoverIsDisplayedBuilder {
	b.properties[`runtimeVersion`] = runtimeVersion

	return b
}

func (b *issueHoverIsDisplayedBuilder) Severity(severity IssueHoverIsDisplayedSeverity) IssueHoverIsDisplayedBuilder {
	b.properties[`severity`] = severity

	return b
}

func (b *issueHoverIsDisplayedBuilder) Build() IssueHoverIsDisplayedEvent {
	return &issueHoverIsDisplayedEvent{
		newBaseEvent(`Issue Hover Is Displayed`, b.properties),
	}
}

type PluginIsInstalledIde string

var PluginIsInstalled = struct {
	Ide struct {
		VisualStudioCode PluginIsInstalledIde

		VisualStudio PluginIsInstalledIde

		Eclipse PluginIsInstalledIde

		JetBrains PluginIsInstalledIde

		Other PluginIsInstalledIde
	}
	Builder func() interface {
		Ide(ide PluginIsInstalledIde) PluginIsInstalledBuilder
	}
}{
	Ide: struct {
		VisualStudioCode PluginIsInstalledIde

		VisualStudio PluginIsInstalledIde

		Eclipse PluginIsInstalledIde

		JetBrains PluginIsInstalledIde

		Other PluginIsInstalledIde
	}{
		VisualStudioCode: `Visual Studio Code`,

		VisualStudio: `Visual Studio`,

		Eclipse: `Eclipse`,

		JetBrains: `JetBrains`,

		Other: `Other`,
	},
	Builder: func() interface {
		Ide(ide PluginIsInstalledIde) PluginIsInstalledBuilder
	} {
		return &pluginIsInstalledBuilder{
			properties: map[string]interface{}{
				`itly`: true,
			},
		}
	},
}

type PluginIsInstalledEvent interface {
	Event
	pluginIsInstalled()
}

type pluginIsInstalledEvent struct {
	baseEvent
}

func (e pluginIsInstalledEvent) pluginIsInstalled() {
}

type PluginIsInstalledBuilder interface {
	Build() PluginIsInstalledEvent
	OsArch(osArch string) PluginIsInstalledBuilder
	OsPlatform(osPlatform string) PluginIsInstalledBuilder
	RuntimeName(runtimeName string) PluginIsInstalledBuilder
	RuntimeVersion(runtimeVersion string) PluginIsInstalledBuilder
}

type pluginIsInstalledBuilder struct {
	properties map[string]interface{}
}

func (b *pluginIsInstalledBuilder) Ide(ide PluginIsInstalledIde) PluginIsInstalledBuilder {
	b.properties[`ide`] = ide

	return b
}

func (b *pluginIsInstalledBuilder) OsArch(osArch string) PluginIsInstalledBuilder {
	b.properties[`osArch`] = osArch

	return b
}

func (b *pluginIsInstalledBuilder) OsPlatform(osPlatform string) PluginIsInstalledBuilder {
	b.properties[`osPlatform`] = osPlatform

	return b
}

func (b *pluginIsInstalledBuilder) RuntimeName(runtimeName string) PluginIsInstalledBuilder {
	b.properties[`runtimeName`] = runtimeName

	return b
}

func (b *pluginIsInstalledBuilder) RuntimeVersion(runtimeVersion string) PluginIsInstalledBuilder {
	b.properties[`runtimeVersion`] = runtimeVersion

	return b
}

func (b *pluginIsInstalledBuilder) Build() PluginIsInstalledEvent {
	return &pluginIsInstalledEvent{
		newBaseEvent(`Plugin Is Installed`, b.properties),
	}
}

type ScanModeIsSelectedEventSource string

type ScanModeIsSelectedIde string

type ScanModeIsSelectedScanMode string

var ScanModeIsSelected = struct {
	EventSource struct {
		Advisor ScanModeIsSelectedEventSource

		App ScanModeIsSelectedEventSource

		Learn ScanModeIsSelectedEventSource

		Ide ScanModeIsSelectedEventSource

		Website ScanModeIsSelectedEventSource

		CodeSnippets ScanModeIsSelectedEventSource
	}
	Ide struct {
		VisualStudioCode ScanModeIsSelectedIde

		VisualStudio ScanModeIsSelectedIde

		Eclipse ScanModeIsSelectedIde

		JetBrains ScanModeIsSelectedIde

		Other ScanModeIsSelectedIde
	}
	ScanMode struct {
		Paused ScanModeIsSelectedScanMode

		Auto ScanModeIsSelectedScanMode

		Manual ScanModeIsSelectedScanMode

		Throttled ScanModeIsSelectedScanMode
	}
	Builder func() interface {
		Ide(ide ScanModeIsSelectedIde) interface {
			ScanMode(scanMode ScanModeIsSelectedScanMode) ScanModeIsSelectedBuilder
		}
	}
}{
	EventSource: struct {
		Advisor ScanModeIsSelectedEventSource

		App ScanModeIsSelectedEventSource

		Learn ScanModeIsSelectedEventSource

		Ide ScanModeIsSelectedEventSource

		Website ScanModeIsSelectedEventSource

		CodeSnippets ScanModeIsSelectedEventSource
	}{
		Advisor: `Advisor`,

		App: `App`,

		Learn: `Learn`,

		Ide: `IDE`,

		Website: `Website`,

		CodeSnippets: `CodeSnippets`,
	},
	Ide: struct {
		VisualStudioCode ScanModeIsSelectedIde

		VisualStudio ScanModeIsSelectedIde

		Eclipse ScanModeIsSelectedIde

		JetBrains ScanModeIsSelectedIde

		Other ScanModeIsSelectedIde
	}{
		VisualStudioCode: `Visual Studio Code`,

		VisualStudio: `Visual Studio`,

		Eclipse: `Eclipse`,

		JetBrains: `JetBrains`,

		Other: `Other`,
	},
	ScanMode: struct {
		Paused ScanModeIsSelectedScanMode

		Auto ScanModeIsSelectedScanMode

		Manual ScanModeIsSelectedScanMode

		Throttled ScanModeIsSelectedScanMode
	}{
		Paused: `paused`,

		Auto: `auto`,

		Manual: `manual`,

		Throttled: `throttled`,
	},
	Builder: func() interface {
		Ide(ide ScanModeIsSelectedIde) interface {
			ScanMode(scanMode ScanModeIsSelectedScanMode) ScanModeIsSelectedBuilder
		}
	} {
		return &scanModeIsSelectedBuilder{
			properties: map[string]interface{}{
				`itly`: true,
			},
		}
	},
}

type ScanModeIsSelectedEvent interface {
	Event
	scanModeIsSelected()
}

type scanModeIsSelectedEvent struct {
	baseEvent
}

func (e scanModeIsSelectedEvent) scanModeIsSelected() {
}

type ScanModeIsSelectedBuilder interface {
	Build() ScanModeIsSelectedEvent
	EventSource(eventSource ScanModeIsSelectedEventSource) ScanModeIsSelectedBuilder
	OsArch(osArch string) ScanModeIsSelectedBuilder
	OsPlatform(osPlatform string) ScanModeIsSelectedBuilder
	RuntimeName(runtimeName string) ScanModeIsSelectedBuilder
	RuntimeVersion(runtimeVersion string) ScanModeIsSelectedBuilder
}

type scanModeIsSelectedBuilder struct {
	properties map[string]interface{}
}

func (b *scanModeIsSelectedBuilder) Ide(ide ScanModeIsSelectedIde) interface {
	ScanMode(scanMode ScanModeIsSelectedScanMode) ScanModeIsSelectedBuilder
} {
	b.properties[`ide`] = ide

	return b
}

func (b *scanModeIsSelectedBuilder) ScanMode(scanMode ScanModeIsSelectedScanMode) ScanModeIsSelectedBuilder {
	b.properties[`scanMode`] = scanMode

	return b
}

func (b *scanModeIsSelectedBuilder) EventSource(eventSource ScanModeIsSelectedEventSource) ScanModeIsSelectedBuilder {
	b.properties[`eventSource`] = eventSource

	return b
}

func (b *scanModeIsSelectedBuilder) OsArch(osArch string) ScanModeIsSelectedBuilder {
	b.properties[`osArch`] = osArch

	return b
}

func (b *scanModeIsSelectedBuilder) OsPlatform(osPlatform string) ScanModeIsSelectedBuilder {
	b.properties[`osPlatform`] = osPlatform

	return b
}

func (b *scanModeIsSelectedBuilder) RuntimeName(runtimeName string) ScanModeIsSelectedBuilder {
	b.properties[`runtimeName`] = runtimeName

	return b
}

func (b *scanModeIsSelectedBuilder) RuntimeVersion(runtimeVersion string) ScanModeIsSelectedBuilder {
	b.properties[`runtimeVersion`] = runtimeVersion

	return b
}

func (b *scanModeIsSelectedBuilder) Build() ScanModeIsSelectedEvent {
	return &scanModeIsSelectedEvent{
		newBaseEvent(`Scan Mode Is Selected`, b.properties),
	}
}

type Ampli struct {
	Disabled bool
	Client   amplitude.Client
	mutex    sync.RWMutex
}

// Load initializes the Ampli wrapper.
// Call once when your application starts.
func (a *Ampli) Load(options LoadOptions) {
	if a.Client != nil {
		log.Print("Warn: Ampli is already initialized. Ampli.Load() should be called once at application start up.")

		return
	}

	var apiKey string
	switch {
	case options.Client.APIKey != "":
		apiKey = options.Client.APIKey
	case options.Environment != "":
		apiKey = APIKey[options.Environment]
	default:
		apiKey = options.Client.Configuration.APIKey
	}

	if apiKey == "" && options.Client.Instance == nil {
		log.Print("Error: Ampli.Load() requires option.Environment, " +
			"and apiKey from either options.Instance.APIKey or APIKey[options.Environment], " +
			"or options.Instance.Instance")
	}

	clientConfig := options.Client.Configuration

	if clientConfig.Plan == nil {
		clientConfig.Plan = &amplitude.Plan{
			Branch:    `main`,
			Source:    `language-server`,
			Version:   `247`,
			VersionID: `9ff98e1b-481c-4d72-a7d8-37de6e04cd11`,
		}
	}

	if clientConfig.IngestionMetadata == nil {
		clientConfig.IngestionMetadata = &amplitude.IngestionMetadata{
			SourceName:    `go-go-ampli`,
			SourceVersion: `2.0.0`,
		}
	}

	if options.Client.Instance != nil {
		a.Client = options.Client.Instance
	} else {
		clientConfig.APIKey = apiKey
		a.Client = amplitude.NewClient(clientConfig)
	}

	a.mutex.Lock()
	a.Disabled = options.Disabled
	a.mutex.Unlock()
}

// InitializedAndEnabled checks if Ampli is initialized and enabled.
func (a *Ampli) InitializedAndEnabled() bool {
	if a.Client == nil {
		log.Print("Error: Ampli is not yet initialized. Have you called Ampli.Load() on app start?")

		return false
	}

	a.mutex.RLock()
	defer a.mutex.RUnlock()

	return !a.Disabled
}

func (a *Ampli) setUserID(userID string, eventOptions *EventOptions) {
	if userID != "" {
		eventOptions.UserID = userID
	}
}

// Track tracks an event.
func (a *Ampli) Track(userID string, event Event, eventOptions ...EventOptions) {
	if !a.InitializedAndEnabled() {
		return
	}

	var options EventOptions
	if len(eventOptions) > 0 {
		options = eventOptions[0]
	}

	a.setUserID(userID, &options)

	baseEvent := event.ToAmplitudeEvent()
	baseEvent.EventOptions = options

	a.Client.Track(baseEvent)
}

// Identify identifies a user and set user properties.
func (a *Ampli) Identify(userID string, identify IdentifyEvent, eventOptions ...EventOptions) {
	a.Track(userID, identify, eventOptions...)
}

// GroupIdentify identifies a group and set group properties.
func (a *Ampli) GroupIdentify(groupType string, groupName string, group GroupEvent, eventOptions ...EventOptions) {
	event := group.ToAmplitudeEvent()
	event.Groups = map[string][]string{groupType: {groupName}}
	if len(eventOptions) > 0 {
		event.EventOptions = eventOptions[0]
	}

	a.Client.Track(event)
}

// SetGroup sets group for the current user.
func (a *Ampli) SetGroup(userID string, groupType string, groupName []string, eventOptions ...EventOptions) {
	var options EventOptions
	if len(eventOptions) > 0 {
		options = eventOptions[0]
	}
	a.setUserID(userID, &options)

	a.Client.SetGroup(groupType, groupName, options)
}

// Flush flushes events waiting in buffer.
func (a *Ampli) Flush() {
	if !a.InitializedAndEnabled() {
		return
	}

	a.Client.Flush()
}

// Shutdown disables and shutdowns Ampli Instance.
func (a *Ampli) Shutdown() {
	if !a.InitializedAndEnabled() {
		return
	}

	a.mutex.Lock()
	a.Disabled = true
	a.mutex.Unlock()

	a.Client.Shutdown()
}

func (a *Ampli) AnalysisIsReady(userID string, event AnalysisIsReadyEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) AnalysisIsTriggered(userID string, event AnalysisIsTriggeredEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) IssueHoverIsDisplayed(userID string, event IssueHoverIsDisplayedEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) PluginIsInstalled(userID string, event PluginIsInstalledEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) ScanModeIsSelected(userID string, event ScanModeIsSelectedEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}
