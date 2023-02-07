package myplugin

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/go-openapi/runtime/middleware"
	"github.com/massalabs/thyra/api/swagger/server/models"
	"github.com/massalabs/thyra/api/swagger/server/restapi/operations"
	"github.com/massalabs/thyra/pkg/plugin"
)

func newExecute(manager *plugin.Manager) operations.PluginManagerExecuteCommandHandler {
	return &execute{manager: manager}
}

type execute struct {
	manager *plugin.Manager
}

//nolint:cyclop
func (e *execute) Handle(params operations.PluginManagerExecuteCommandParams) middleware.Responder {
	cmd := params.Body.Command

	log.Printf("[POST /plugin-manager/%d/execute] command: %s", params.ID, cmd)

	plugin, err := e.manager.Plugin(params.ID)
	if err != nil {
		return operations.NewPluginManagerExecuteCommandNotFound().WithPayload(
			&models.Error{Code: errorCodePluginUnknown, Message: fmt.Sprintf("get plugin error: %s", err.Error())})
	}

	status := plugin.Status()

	pluginName := filepath.Base(plugin.BinPath)

	switch cmd {
	case "start":
		err := plugin.Start()
		if err != nil {
			return executeFailed(cmd, status,
				fmt.Sprintf("Error while starting plugin %s: %s.\n", pluginName, err))
		}
	case "stop":
		err := plugin.Stop()
		if err != nil {
			return executeFailed(cmd, status, fmt.Sprintf("Error while stopping plugin %s: %s.\n", pluginName, err))
		}
	case "restart":
		err := plugin.Stop()
		if err != nil {
			return executeFailed(cmd, status, fmt.Sprintf("Error while stopping plugin %s: %s.\n", pluginName, err))
		}

		err = plugin.Start()
		if err != nil {
			return executeFailed(cmd, status,
				fmt.Sprintf("Error while restarting plugin %s: %s.\n", pluginName, err))
		}
	case "update":
		return operations.NewPluginManagerExecuteCommandNotImplemented()
	default:
		return executeFailed(cmd, status, fmt.Sprintf("Unknown command %s.\n", cmd))
	}

	return operations.NewPluginManagerExecuteCommandNoContent()
}

func executeFailed(cmd string, currentStatus plugin.Status, errorMsg string,
) *operations.PluginManagerExecuteCommandBadRequest {
	errStr := ""
	if errorMsg != "" {
		errStr = fmt.Sprintf(" (%s)", errorMsg)
	}

	return operations.NewPluginManagerExecuteCommandBadRequest().WithPayload(
		&models.Error{
			Code:    errorCodePluginExecuteCmdBadRequest,
			Message: fmt.Sprintf("[%s] %s. Current plugin status is %s.", cmd, errStr, currentStatus),
		})
}
