//go:build libvirt && cgo

// (C) Copyright Confidential Containers Contributors
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"os"
	"testing"

	"github.com/confidential-containers/cloud-api-adaptor/src/cloud-api-adaptor/test/provisioner/libvirt"
)

func TestLibvirtCreateSimplePod(t *testing.T) {
	assert := LibvirtAssert{}
	DoTestCreateSimplePod(t, testEnv, assert)
}

func TestLibvirtCreatePodWithConfigMap(t *testing.T) {
	SkipTestOnCI(t)
	assert := LibvirtAssert{}
	DoTestCreatePodWithConfigMap(t, testEnv, assert)
}

func TestLibvirtCreatePodWithSecret(t *testing.T) {
	assert := LibvirtAssert{}
	DoTestCreatePodWithSecret(t, testEnv, assert)
}

func TestLibvirtCreatePeerPodContainerWithExternalIPAccess(t *testing.T) {
	SkipTestOnCI(t)
	assert := LibvirtAssert{}
	DoTestCreatePeerPodContainerWithExternalIPAccess(t, testEnv, assert)

}

func TestLibvirtCreatePeerPodContainerWithValidAlternateImage(t *testing.T) {
	assert := LibvirtAssert{}
	DoTestCreatePeerPodContainerWithValidAlternateImage(t, testEnv, assert, libvirt.AlternateVolumeName)
}

func TestLibvirtCreatePeerPodContainerWithInvalidAlternateImage(t *testing.T) {
	assert := LibvirtAssert{}
	nonExistingImageName := "non-existing-image"
	expectedErrorMessage := "Error in creating volume: Can't retrieve volume " + nonExistingImageName
	DoTestCreatePeerPodContainerWithInvalidAlternateImage(t, testEnv, assert, nonExistingImageName, expectedErrorMessage)
}

func TestLibvirtCreatePeerPodWithJob(t *testing.T) {
	assert := LibvirtAssert{}
	DoTestCreatePeerPodWithJob(t, testEnv, assert)
}

func TestLibvirtCreatePeerPodAndCheckUserLogs(t *testing.T) {
	assert := LibvirtAssert{}
	DoTestCreatePeerPodAndCheckUserLogs(t, testEnv, assert)
}

func TestLibvirtCreatePeerPodAndCheckWorkDirLogs(t *testing.T) {
	// This test is causing issues on CI with instability, so skip until we can resolve this.
	// See https://github.com/confidential-containers/cloud-api-adaptor/issues/1831
	SkipTestOnCI(t)
	assert := LibvirtAssert{}
	DoTestCreatePeerPodAndCheckWorkDirLogs(t, testEnv, assert)
}

func TestLibvirtCreatePeerPodAndCheckEnvVariableLogsWithImageOnly(t *testing.T) {
	// This test is causing issues on CI with instability, so skip until we can resolve this.
	// See https://github.com/confidential-containers/cloud-api-adaptor/issues/1831
	SkipTestOnCI(t)
	assert := LibvirtAssert{}
	DoTestCreatePeerPodAndCheckEnvVariableLogsWithImageOnly(t, testEnv, assert)
}

func TestLibvirtCreatePeerPodAndCheckEnvVariableLogsWithDeploymentOnly(t *testing.T) {
	// This test is causing issues on CI with instability, so skip until we can resolve this.
	// See https://github.com/confidential-containers/cloud-api-adaptor/issues/1831
	SkipTestOnCI(t)
	assert := LibvirtAssert{}
	DoTestCreatePeerPodAndCheckEnvVariableLogsWithDeploymentOnly(t, testEnv, assert)
}

func TestLibvirtCreatePeerPodAndCheckEnvVariableLogsWithImageAndDeployment(t *testing.T) {
	// This test is causing issues on CI with instability, so skip until we can resolve this.
	// See https://github.com/confidential-containers/cloud-api-adaptor/issues/1831
	SkipTestOnCI(t)
	assert := LibvirtAssert{}
	DoTestCreatePeerPodAndCheckEnvVariableLogsWithImageAndDeployment(t, testEnv, assert)
}

func TestLibvirtCreateNginxDeployment(t *testing.T) {
	// This test is causing issues on CI with instability, so skip until we can resolve this.
	// See https://github.com/confidential-containers/cloud-api-adaptor/issues/2046
	SkipTestOnCI(t)
	assert := LibvirtAssert{}
	DoTestNginxDeployment(t, testEnv, assert)
}

/*
Failing due to issues will pulling image (ErrImagePull)
func TestLibvirtCreatePeerPodWithLargeImage(t *testing.T) {
	assert := LibvirtAssert{}
	DoTestCreatePeerPodWithLargeImage(t, testEnv, assert)
}
*/

func TestLibvirtDeletePod(t *testing.T) {
	assert := LibvirtAssert{}
	DoTestDeleteSimplePod(t, testEnv, assert)
}

func TestLibvirtPodToServiceCommunication(t *testing.T) {
	// This test is causing issues on CI with instability, so skip until we can resolve this.
	SkipTestOnCI(t)
	assert := LibvirtAssert{}
	DoTestPodToServiceCommunication(t, testEnv, assert)
}

func TestLibvirtPodsMTLSCommunication(t *testing.T) {
	// This test is causing issues on CI with instability, so skip until we can resolve this.
	SkipTestOnCI(t)
	assert := LibvirtAssert{}
	DoTestPodsMTLSCommunication(t, testEnv, assert)
}

func TestLibvirtSealedSecret(t *testing.T) {
	if !isTestWithKbs() {
		t.Skip("Skipping kbs related test as kbs is not deployed")
	}
	_ = keyBrokerService.SetSampleSecretKey()
	_ = keyBrokerService.EnableKbsCustomizedResourcePolicy("allow_all.rego")
	_ = keyBrokerService.EnableKbsCustomizedAttestationPolicy("allow_all.rego")
	kbsEndpoint, _ := keyBrokerService.GetCachedKbsEndpoint()
	assert := LibvirtAssert{}
	DoTestSealedSecret(t, testEnv, assert, kbsEndpoint)
}

func TestLibvirtKbsKeyRelease(t *testing.T) {
	if !isTestWithKbs() {
		t.Skip("Skipping kbs related test as kbs is not deployed")
	}
	_ = keyBrokerService.SetSampleSecretKey()
	_ = keyBrokerService.EnableKbsCustomizedResourcePolicy("allow_all.rego")
	_ = keyBrokerService.EnableKbsCustomizedAttestationPolicy("deny_all.rego")
	kbsEndpoint, _ := keyBrokerService.GetCachedKbsEndpoint()
	assert := LibvirtAssert{}
	t.Parallel()
	DoTestKbsKeyReleaseForFailure(t, testEnv, assert, kbsEndpoint)
	if isTestWithKbsIBMSE() {
		t.Log("KBS with ibmse cases")
		// the allow_*_.rego file is created by follow document
		// https://github.com/confidential-containers/trustee/blob/main/deps/verifier/src/se/README.md#set-attestation-policy
		_ = keyBrokerService.EnableKbsCustomizedAttestationPolicy("allow_with_wrong_image_tag.rego")
		DoTestKbsKeyReleaseForFailure(t, testEnv, assert, kbsEndpoint)
		_ = keyBrokerService.EnableKbsCustomizedAttestationPolicy("allow_with_correct_claims.rego")
		DoTestKbsKeyRelease(t, testEnv, assert, kbsEndpoint)
	} else {
		t.Log("KBS normal cases")
		_ = keyBrokerService.EnableKbsCustomizedAttestationPolicy("allow_all.rego")
		DoTestKbsKeyRelease(t, testEnv, assert, kbsEndpoint)
	}
}

func TestLibvirtRestrictivePolicyBlocksExec(t *testing.T) {
	assert := LibvirtAssert{}
	DoTestRestrictivePolicyBlocksExec(t, testEnv, assert)
}

func TestLibvirtPermissivePolicyAllowsExec(t *testing.T) {
	assert := LibvirtAssert{}
	DoTestPermissivePolicyAllowsExec(t, testEnv, assert)
}

func TestLibvirtCreatePeerPodWithAuthenticatedImageWithoutCredentials(t *testing.T) {
	assert := LibvirtAssert{}
	if os.Getenv("AUTHENTICATED_REGISTRY_IMAGE") != "" {
		DoTestCreatePeerPodWithAuthenticatedImageWithoutCredentials(t, testEnv, assert)
	} else {
		t.Skip("Authenticated Image Name not exported")
	}
}

func TestLibvirtCreatePeerPodWithAuthenticatedImageWithValidCredentials(t *testing.T) {
	assert := LibvirtAssert{}
	if os.Getenv("REGISTRY_CREDENTIAL_ENCODED") != "" && os.Getenv("AUTHENTICATED_REGISTRY_IMAGE") != "" {
		DoTestCreatePeerPodWithAuthenticatedImageWithValidCredentials(t, testEnv, assert)
	} else {
		t.Skip("Registry Credentials, or authenticated image name not exported")
	}
}

func TestLibvirtCreateWithCpuAndMemRequestLimit(t *testing.T) {
	assert := LibvirtAssert{}
	DoTestPodWithCpuMemLimitsAndRequests(t, testEnv, assert, "100m", "100Mi", "200m", "200Mi")
}
