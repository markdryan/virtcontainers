//
// Copyright (c) 2018 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package virtcontainers

import (
	"fmt"
	"plugin"
)

const (
	initFuncName                = "Init"
	createPodFuncName           = "CreatePod"
	startPodFuncName            = "StartPod"
	waitPodFuncName             = "WaitPod"
	stopPodFuncName             = "StopPod"
	pausePodFuncName            = "PausePod"
	resumePodFuncName           = "ResumePod"
	addDeviceFuncName           = "AddDevice"
	hotplugAddDeviceFuncName    = "HotplugAddDevice"
	hotplugRemoveDeviceFuncName = "HotplugRemoveDevice"
	getPodConsoleFuncName       = "GetPodConsole"
	capabilitiesFuncName        = "Capabilities"
)

type pluginHypervisor struct {
	handler    *plugin.Plugin
	hypervisor interface{}
}

func (h *pluginHypervisor) init(pod *Pod) (err error) {
	pluginPath := pod.config.HypervisorConfig.PluginPath

	h.handler, err = plugin.Open(pluginPath)
	if err != nil {
		return fmt.Errorf("Failed to open plugin path %q: %v",
			pluginPath, err)
	}

	initFunc, err := h.handler.Lookup(initFuncName)
	if err != nil {
		return fmt.Errorf("Failed to lookup function %q: %v",
			initFuncName, err)
	}

	h.hypervisor, err = initFunc.(func(*Pod) (interface{}, error))(pod)
	if err != nil {
		return fmt.Errorf("%s() failed: %v", initFuncName, err)
	}

	return nil
}

func (h *pluginHypervisor) createPod(podConfig PodConfig) error {
	createPodFunc, err := h.handler.Lookup(createPodFuncName)
	if err != nil {
		return fmt.Errorf("Failed to lookup function %q: %v",
			createPodFuncName, err)
	}

	if err := createPodFunc.(func(interface{}, PodConfig) error)(h.hypervisor, podConfig); err != nil {
		return fmt.Errorf("%s() failed: %v", createPodFuncName, err)
	}

	return nil
}

func (h *pluginHypervisor) startPod() error {
	startPodFunc, err := h.handler.Lookup(startPodFuncName)
	if err != nil {
		return fmt.Errorf("Failed to lookup function %q: %v",
			startPodFuncName, err)
	}

	if err := startPodFunc.(func(interface{}) error)(h.hypervisor); err != nil {
		return fmt.Errorf("%s() failed: %v", startPodFuncName, err)
	}

	return nil
}

func (h *pluginHypervisor) waitPod(timeout int) error {
	waitPodFunc, err := h.handler.Lookup(waitPodFuncName)
	if err != nil {
		return fmt.Errorf("Failed to lookup function %q: %v",
			waitPodFuncName, err)
	}

	if err := waitPodFunc.(func(interface{}, int) error)(h.hypervisor, timeout); err != nil {
		return fmt.Errorf("%s() failed: %v", waitPodFuncName, err)
	}

	return nil
}

func (h *pluginHypervisor) stopPod() error {
	stopPodFunc, err := h.handler.Lookup(stopPodFuncName)
	if err != nil {
		return fmt.Errorf("Failed to lookup function %q: %v",
			stopPodFuncName, err)
	}

	if err := stopPodFunc.(func(interface{}) error)(h.hypervisor); err != nil {
		return fmt.Errorf("%s() failed: %v", stopPodFuncName, err)
	}

	return nil
}

func (h *pluginHypervisor) pausePod() error {
	pausePodFunc, err := h.handler.Lookup(pausePodFuncName)
	if err != nil {
		return fmt.Errorf("Failed to lookup function %q: %v",
			pausePodFuncName, err)
	}

	if err := pausePodFunc.(func(interface{}) error)(h.hypervisor); err != nil {
		return fmt.Errorf("%s() failed: %v", pausePodFuncName, err)
	}

	return nil
}

func (h *pluginHypervisor) resumePod() error {
	resumePodFunc, err := h.handler.Lookup(resumePodFuncName)
	if err != nil {
		return fmt.Errorf("Failed to lookup function %q: %v",
			resumePodFuncName, err)
	}

	if err := resumePodFunc.(func(interface{}) error)(h.hypervisor); err != nil {
		return fmt.Errorf("%s() failed: %v", resumePodFuncName, err)
	}

	return nil
}

func (h *pluginHypervisor) addDevice(devInfo interface{}, devType DeviceType) error {
	addDeviceFunc, err := h.handler.Lookup(addDeviceFuncName)
	if err != nil {
		return fmt.Errorf("Failed to lookup function %q: %v",
			addDeviceFuncName, err)
	}

	if err := addDeviceFunc.(func(interface{}, interface{}, DeviceType) error)(h.hypervisor, devInfo, devType); err != nil {
		return fmt.Errorf("%s() failed: %v", addDeviceFuncName, err)
	}

	return nil
}

func (h *pluginHypervisor) hotplugAddDevice(devInfo interface{}, devType DeviceType) error {
	hotplugAddDeviceFunc, err := h.handler.Lookup(hotplugAddDeviceFuncName)
	if err != nil {
		return fmt.Errorf("Failed to lookup function %q: %v",
			hotplugAddDeviceFuncName, err)
	}

	if err := hotplugAddDeviceFunc.(func(interface{}, interface{}, DeviceType) error)(h.hypervisor, devInfo, devType); err != nil {
		return fmt.Errorf("%s() failed: %v", hotplugAddDeviceFuncName, err)
	}

	return nil
}

func (h *pluginHypervisor) hotplugRemoveDevice(devInfo interface{}, devType DeviceType) error {
	hotplugRemoveDeviceFunc, err := h.handler.Lookup(hotplugRemoveDeviceFuncName)
	if err != nil {
		return fmt.Errorf("Failed to lookup function %q: %v",
			hotplugRemoveDeviceFuncName, err)
	}

	if err := hotplugRemoveDeviceFunc.(func(interface{}, interface{}, DeviceType) error)(h.hypervisor, devInfo, devType); err != nil {
		return fmt.Errorf("%s() failed: %v", hotplugRemoveDeviceFuncName, err)
	}

	return nil
}

func (h *pluginHypervisor) getPodConsole(podID string) string {
	getPodConsoleFunc, err := h.handler.Lookup(getPodConsoleFuncName)
	if err != nil {
		return ""
	}

	return getPodConsoleFunc.(func(interface{}, string) string)(h.hypervisor, podID)
}

func (h *pluginHypervisor) capabilities() Capabilities {
	capabilitiesFunc, err := h.handler.Lookup(capabilitiesFuncName)
	if err != nil {
		return Capabilities{}
	}

	return capabilitiesFunc.(func(interface{}) Capabilities)(h.hypervisor)
}
