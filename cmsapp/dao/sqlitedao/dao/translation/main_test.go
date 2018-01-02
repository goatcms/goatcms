package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Translation {
	var (
		key   string = "OZYIvQREApDxw629AoGPzd1VTDMhuAu27YWJNbA1AyRB"
		value string = "zLsuzTgVFSQQjYhUQirmdzIgrVmDS7Zii4Hw3VoabQvng0UR1T9QydcPdAOA3dxkLsdTJzgHGJlTlafhjw2JtW7nHljz9W9LB1xUYO3FrnD8J6W3cSJCUKM0qudt8zUgzq5RLPJSo0VNqJQBLmiOuwE5ko7Ze8meHSogNx7g0Qf7CX2hGSmIvNCcqvPKWfjgYrWL5MqIof0OOmGJkfv8sZnsIeG 02sUkqG27LgeDQ8PPdkNiUvfYDfcHBIOqjNekwcJRZz3K6wEy90rThPR XU7NaXxt6YpPwCsToiXDE5D62AbUIZLlHw0j8QhgmVwZStIe1iTIV4btfWhYzjmZ8qCLWKCyg759GEgHQ2jVMZKQr6jHY2PEK8mV2F0Ja7XmI6VpU3NfHLCaxNa"
	)
	return &entities.Translation{
		Key:   &key,
		Value: &value,
	}
}

func NewMockEntity2() *entities.Translation {
	var (
		value string = "UZ05Ehr  F9B7w5aRvErBWp7eujlvuofcyy7DsSciTBlh RflG8duJNnevdsMagRlbv EQCvtLsXrKA53W6fddWCmtD4DOrptW9Ur5owkFL4o9oq69fY4fRH1Vt7EcoeUi4TeL2fKdaBWkl0EtkePw1bRaykpasYh9augqfySLSpQZ7WYJ AxIRl7PbyMrkLb9Yxx4ESO xJL07ZL7IEBa9vz9JvxeKBDsdA1nIgLV779Fg3JyjsdK7xVAIrSQbDIdgTepNpaLto4Vl2wQiUTXsfQctZ2TyT4YkZRg1eCG6qoiv9V5b JvmXg fPeAEHnIxAQZn2H4KydxxT7UB9fFeifmpEnvtBajKPW4RjbVu4wBFk7rqwcrBvcWH6vqK7d3DvwK8fQzscskg0"
		key   string = "lhMPBTczc6RGrjNgMJwZTuyQvB22PCzXAUQvpsCtTs01"
	)
	return &entities.Translation{
		Value: &value,
		Key:   &key,
	}
}
