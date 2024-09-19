// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package vm

import (
	"csbench/config"
	"csbench/utils"
	"log"

	"github.com/sjyu1/ablestack-mold-go/v2/cloudstack"
)

func ListVMs(cs *cloudstack.CloudStackClient, domainId string) ([]*cloudstack.VirtualMachine, error) {
	result := make([]*cloudstack.VirtualMachine, 0)
	page := 1
	p := cs.VirtualMachine.NewListVirtualMachinesParams()
	p.SetDomainid(domainId)
	p.SetPagesize(config.PageSize)
	for {
		p.SetPage(page)
		resp, err := cs.VirtualMachine.ListVirtualMachines(p)
		if err != nil {
			log.Printf("Failed to list vm due to %v", err)
			return result, err
		}
		result = append(result, resp.VirtualMachines...)
		if len(result) < resp.Count {
			page++
		} else {
			break
		}
	}
	return result, nil
}

func DeployVm(cs *cloudstack.CloudStackClient, domainId string, networkId string, account string) (*cloudstack.DeployVirtualMachineResponse, error) {
	vmName := "Vm-" + utils.RandomString(10)
	p := cs.VirtualMachine.NewDeployVirtualMachineParams(config.ServiceOfferingId, config.TemplateId, vmName)
	p.SetDomainid(domainId)
	p.SetZoneid(config.ZoneId)
	p.SetNetworkids([]string{networkId})
	p.SetName(vmName)
	p.SetAccount(account)
	p.SetStartvm(config.StartVM)

	resp, err := cs.VirtualMachine.DeployVirtualMachine(p)
	if err != nil {
		log.Printf("Failed to deploy vm due to: %v", err)
		return nil, err
	}
	return resp, nil
}

func DestroyVm(cs *cloudstack.CloudStackClient, vmId string) error {

	deleteParams := cs.VirtualMachine.NewDestroyVirtualMachineParams(vmId)
	deleteParams.SetExpunge(true)
	_, err := cs.VirtualMachine.DestroyVirtualMachine(deleteParams)
	if err != nil {
		log.Printf("Failed to destroy Vm with Id %s due to %v", vmId, err)
		return err
	}
	return nil
}

func StartVM(cs *cloudstack.CloudStackClient, vmId string) error {
	p := cs.VirtualMachine.NewStartVirtualMachineParams(vmId)
	_, err := cs.VirtualMachine.StartVirtualMachine(p)
	if err != nil {
		log.Printf("Failed to start vm with id %s due to %v", vmId, err)
		return err
	}
	return nil
}

func StopVM(cs *cloudstack.CloudStackClient, vmId string) error {
	p := cs.VirtualMachine.NewStopVirtualMachineParams(vmId)
	_, err := cs.VirtualMachine.StopVirtualMachine(p)
	if err != nil {
		log.Printf("Failed to stop vm with id %s due to %v", vmId, err)
		return err
	}
	return nil
}

func RebootVM(cs *cloudstack.CloudStackClient, vmId string) error {
	p := cs.VirtualMachine.NewRebootVirtualMachineParams(vmId)
	_, err := cs.VirtualMachine.RebootVirtualMachine(p)
	if err != nil {
		log.Printf("Failed to reboot vm with id %s due to %v", vmId, err)
		return err
	}
	return nil
}

func CreateVMSnapshot(cs *cloudstack.CloudStackClient, vmId string) (*cloudstack.CreateVMSnapshotResponse, error) {
	p := cs.Snapshot.NewCreateVMSnapshotParams(vmId)
	resp, err := cs.Snapshot.CreateVMSnapshot(p)
	if err != nil {
		log.Printf("Failed to create vm snapshot with id %s due to %v", vmId, err)
		return nil, err
	}
	return resp, nil
}

func DeleteVMSnapshot(cs *cloudstack.CloudStackClient, snapshotId string) error {
	p := cs.Snapshot.NewDeleteVMSnapshotParams(snapshotId)
	_, err := cs.Snapshot.DeleteVMSnapshot(p)
	if err != nil {
		log.Printf("Failed to delete vm snapshot with id %s due to %v", snapshotId, err)
		return err
	}
	return nil
}

func AllocateVbmcToVM(cs *cloudstack.CloudStackClient, vmId string) error {
	p := cs.VirtualMachine.NewAllocateVbmcToVMParams(vmId)
	_, err := cs.VirtualMachine.AllocateVbmcToVM(p)
	if err != nil {
		log.Printf("Failed to allocate vbmc to vm with id %s due to %v", vmId, err)
		return err
	}
	return nil
}

func RemoveVbmcToVM(cs *cloudstack.CloudStackClient, vmId string) error {
	p := cs.VirtualMachine.NewRemoveVbmcToVMParams(vmId)
	_, err := cs.VirtualMachine.RemoveVbmcToVM(p)
	if err != nil {
		log.Printf("Failed to remove vbmc to vm with id %s due to %v", vmId, err)
		return err
	}
	return nil
}

// func CloneVirtualMachine(cs *cloudstack.CloudStackClient, vmId string) error {
// 	p := cs.VirtualMachine.NewCloneVirtualMachineParams(vmId)
// 	_, err := cs.VirtualMachine.CloneVirtualMachine(p)
// 	if err != nil {
// 		log.Printf("Failed to clone to vm with id %s due to %v", vmId, err)
// 		return err
// 	}
// 	return nil
// }

func RestoreVirtualMachine(cs *cloudstack.CloudStackClient, vmId string) error {
	p := cs.VirtualMachine.NewRestoreVirtualMachineParams(vmId)
	_, err := cs.VirtualMachine.RestoreVirtualMachine(p)
	if err != nil {
		log.Printf("Failed to restore to vm with id %s due to %v", vmId, err)
		return err
	}
	return nil
}
