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

package snyk

import (
	"context"
	"sync"

	"github.com/snyk/go-application-framework/pkg/auth"
)

const (
	NavigateToRangeCommand       = "snyk.navigateToRange"
	WorkspaceScanCommand         = "snyk.workspace.scan"
	WorkspaceFolderScanCommand   = "snyk.workspaceFolder.scan"
	OpenBrowserCommand           = "snyk.openBrowser"
	LoginCommand                 = "snyk.login"
	CopyAuthLinkCommand          = "snyk.copyAuthLink"
	LogoutCommand                = "snyk.logout"
	TrustWorkspaceFoldersCommand = "snyk.trustWorkspaceFolders"
	OAuthRefreshCommand          = "snyk.oauthRefreshCommand"
	OpenLearnLesson              = "snyk.openLearnLesson"
	GetLearnLesson               = "snyk.getLearnLesson"
	GetSettingsSastEnabled       = "snyk.getSettingsSastEnabled"
	GetActiveUser                = "snyk.getActiveUser"
)

var (
	DefaultOpenBrowserFunc = func(url string) { auth.OpenBrowser(url) }
)

type Command interface {
	Command() CommandData
	Execute(ctx context.Context) (any, error)
}

type CommandData struct {
	/**
	 * Title of the command, like `save`.
	 */
	Title string
	/**
	 * The identifier of the actual command handler.
	 */
	CommandId string
	/**
	 * Arguments that the command handler should be
	 * invoked with.
	 */
	Arguments []any
}

type CommandName string

type CommandService interface {
	ExecuteCommand(ctx context.Context, command Command) (any, error)
}

type CommandServiceMock struct {
	m                sync.Mutex
	executedCommands []Command
}

func NewCommandServiceMock() *CommandServiceMock {
	return &CommandServiceMock{}
}

func (service *CommandServiceMock) ExecuteCommand(_ context.Context, command Command) (any, error) {
	service.m.Lock()
	service.executedCommands = append(service.executedCommands, command)
	service.m.Unlock()
	return nil, nil
}
func (service *CommandServiceMock) ExecutedCommands() []Command {
	service.m.Lock()
	cmds := service.executedCommands
	service.m.Unlock()
	return cmds
}
