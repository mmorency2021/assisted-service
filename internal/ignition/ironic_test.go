package ignition

import (
	"net/url"

	config_32 "github.com/coreos/ignition/v2/config/v3_2"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/models"
)

var _ = Describe("Ignition with converged flow", func() {
	var (
		infraEnv      common.InfraEnv
		infraEnvID    strfmt.UUID
		ironicBaseURL string
		iib           IronicIgniotionBuilder
	)
	BeforeEach(func() {
		infraEnvID = "a64fff36-dcb1-11ea-87d0-0242ac130003"
		ironicBaseURL = "https://10.10.10.10"

		infraEnv = common.InfraEnv{
			InfraEnv: models.InfraEnv{
				ID:            &infraEnvID,
				PullSecretSet: false,
			},
			PullSecret: "{\"auths\":{\"cloud.openshift.com\":{\"auth\":\"dG9rZW46dGVzdAo=\",\"email\":\"coyote@acme.com\"}}}",
		}
	})
	It("GenerateIronicConfig success", func() {
		config := IronicIgniotionBuilderConfig{BaremetalIronicAgentImage: "defaultIronicAgentImage:latest", BaremetalIronicAgentImageForArm: "defaultIronicAgentImageForArm:latest"}
		iib = NewIronicIgniotionBuilder(config)

		conf, err := iib.GenerateIronicConfig(ironicBaseURL, infraEnv, "")
		Expect(err).NotTo(HaveOccurred())
		Expect(conf).Should(ContainSubstring("ironic-agent.service"))
		Expect(conf).Should(ContainSubstring(url.QueryEscape(ironicBaseURL)))
		Expect(conf).Should(ContainSubstring(config.BaremetalIronicAgentImage))
		validateIgnition(conf)
	})
	It("GenerateIronicConfig override default ironic agent image", func() {
		iib = NewIronicIgniotionBuilder(IronicIgniotionBuilderConfig{BaremetalIronicAgentImage: "defaultIronicAgentImage:latest"})
		ironicAgentImage := "ironicAgentImage:custom"
		conf, err := iib.GenerateIronicConfig(ironicBaseURL, infraEnv, ironicAgentImage)
		Expect(err).NotTo(HaveOccurred())
		Expect(conf).Should(ContainSubstring("ironic-agent.service"))
		Expect(conf).Should(ContainSubstring(url.QueryEscape(ironicBaseURL)))
		Expect(conf).Should(ContainSubstring(ironicAgentImage))
		validateIgnition(conf)
	})
	It("GenerateIronicConfig use default ironic agent image for arm", func() {
		config := IronicIgniotionBuilderConfig{BaremetalIronicAgentImage: "defaultIronicAgentImage:latest", BaremetalIronicAgentImageForArm: "defaultIronicAgentImageForArm:latest"}
		iib = NewIronicIgniotionBuilder(config)
		infraEnv.CPUArchitecture = common.ARM64CPUArchitecture

		conf, err := iib.GenerateIronicConfig(ironicBaseURL, infraEnv, "")
		Expect(err).NotTo(HaveOccurred())
		Expect(conf).Should(ContainSubstring("ironic-agent.service"))
		Expect(conf).Should(ContainSubstring(url.QueryEscape(ironicBaseURL)))
		Expect(conf).Should(ContainSubstring(config.BaremetalIronicAgentImageForArm))
		Expect(conf).ShouldNot(ContainSubstring(config.BaremetalIronicAgentImage))
		validateIgnition(conf)
	})
	It("GenerateIronicConfig missing ironic service URL", func() {
		iib = NewIronicIgniotionBuilder(IronicIgniotionBuilderConfig{BaremetalIronicAgentImage: "defaultIronicAgentImage:latest"})
		_, err := iib.GenerateIronicConfig("", infraEnv, "")
		Expect(err).To(HaveOccurred())
	})
	It("GenerateIronicConfig missing ironic agent image", func() {
		iib = NewIronicIgniotionBuilder(IronicIgniotionBuilderConfig{BaremetalIronicAgentImage: ""})
		_, err := iib.GenerateIronicConfig(ironicBaseURL, infraEnv, "")
		Expect(err).To(HaveOccurred())
	})
})

func validateIgnition(conf []byte) {
	v32Config, _, err1 := config_32.Parse(conf)
	Expect(err1).ToNot(HaveOccurred())
	Expect(v32Config.Ignition.Version).To(Equal("3.2.0"))
	// We expect a single file and a single systemd unit
	Expect(len(v32Config.Storage.Files)).To(Equal(1))
	Expect(len(v32Config.Systemd.Units)).To(Equal(1))
}
