package schema
	
import (
	"errors"
)

func GetFhirResourceInstance(resource_name string) (interface{}, error) {

	FhirStructMap := map[string]interface{}{
		"MedicationRequestDispenseRequest" : &MedicationRequestDispenseRequest{}, 
		"MedicinalProductPharmaceuticalCharacteristics" : &MedicinalProductPharmaceuticalCharacteristics{}, 
		"TerminologyCapabilitiesExpansion" : &TerminologyCapabilitiesExpansion{}, 
		"ProductShelfLife" : &ProductShelfLife{}, 
		"CarePlanActivity" : &CarePlanActivity{}, 
		"Condition" : &Condition{}, 
		"Linkage" : &Linkage{}, 
		"Library" : &Library{}, 
		"SpecimenContainer" : &SpecimenContainer{}, 
		"SubstanceIngredient" : &SubstanceIngredient{}, 
		"ElementDefinitionBinding" : &ElementDefinitionBinding{}, 
		"CarePlan" : &CarePlan{}, 
		"ExplanationOfBenefitPayee" : &ExplanationOfBenefitPayee{}, 
		"ExplanationOfBenefitDiagnosis" : &ExplanationOfBenefitDiagnosis{}, 
		"NamingSystemUniqueId" : &NamingSystemUniqueId{}, 
		"ValueSetParameter" : &ValueSetParameter{}, 
		"VisionPrescriptionLensSpecification" : &VisionPrescriptionLensSpecification{}, 
		"VisionPrescriptionPrism" : &VisionPrescriptionPrism{}, 
		"ChargeItemDefinitionPriceComponent" : &ChargeItemDefinitionPriceComponent{}, 
		"ClaimDetail" : &ClaimDetail{}, 
		"ClaimResponseProcessNote" : &ClaimResponseProcessNote{}, 
		"DeviceDefinitionUdiDeviceIdentifier" : &DeviceDefinitionUdiDeviceIdentifier{}, 
		"BiologicallyDerivedProductManipulation" : &BiologicallyDerivedProductManipulation{}, 
		"ExplanationOfBenefitSubDetail1" : &ExplanationOfBenefitSubDetail1{}, 
		"LocationHoursOfOperation" : &LocationHoursOfOperation{}, 
		"ValueSetConcept" : &ValueSetConcept{}, 
		"ExampleScenarioAlternative" : &ExampleScenarioAlternative{}, 
		"GoalTarget" : &GoalTarget{}, 
		"CodeableConcept" : &CodeableConcept{}, 
		"ElementDefinition" : &ElementDefinition{}, 
		"CarePlanDetail" : &CarePlanDetail{}, 
		"ClaimInsurance" : &ClaimInsurance{}, 
		"ContractLegal" : &ContractLegal{}, 
		"ImmunizationRecommendationDateCriterion" : &ImmunizationRecommendationDateCriterion{}, 
		"TestScriptAction2" : &TestScriptAction2{}, 
		"Age" : &Age{}, 
		"Dosage" : &Dosage{}, 
		"SubstanceSpecificationRelationship" : &SubstanceSpecificationRelationship{}, 
		"Immunization" : &Immunization{}, 
		"ConceptMap" : &ConceptMap{}, 
		"ExplanationOfBenefitInsurance" : &ExplanationOfBenefitInsurance{}, 
		"ExplanationOfBenefitSubDetail" : &ExplanationOfBenefitSubDetail{}, 
		"GraphDefinitionLink" : &GraphDefinitionLink{}, 
		"Person" : &Person{}, 
		"Practitioner" : &Practitioner{}, 
		"ConditionEvidence" : &ConditionEvidence{}, 
		"ImplementationGuideDependsOn" : &ImplementationGuideDependsOn{}, 
		"InsurancePlanSpecificCost" : &InsurancePlanSpecificCost{}, 
		"MeasureReportPopulation" : &MeasureReportPopulation{}, 
		"MolecularSequenceInner" : &MolecularSequenceInner{}, 
		"ParameterDefinition" : &ParameterDefinition{}, 
		"AdverseEventCausality" : &AdverseEventCausality{}, 
		"CareTeamParticipant" : &CareTeamParticipant{}, 
		"CompositionAttester" : &CompositionAttester{}, 
		"MedicationDispensePerformer" : &MedicationDispensePerformer{}, 
		"MessageHeaderDestination" : &MessageHeaderDestination{}, 
		"TestReportTest" : &TestReportTest{}, 
		"ClaimCareTeam" : &ClaimCareTeam{}, 
		"ClaimResponseAddItem" : &ClaimResponseAddItem{}, 
		"MedicationAdministrationPerformer" : &MedicationAdministrationPerformer{}, 
		"PatientLink" : &PatientLink{}, 
		"SpecimenDefinition" : &SpecimenDefinition{}, 
		"SpecimenDefinitionTypeTested" : &SpecimenDefinitionTypeTested{}, 
		"CoverageEligibilityRequestInsurance" : &CoverageEligibilityRequestInsurance{}, 
		"HealthcareServiceNotAvailable" : &HealthcareServiceNotAvailable{}, 
		"MolecularSequenceStructureVariant" : &MolecularSequenceStructureVariant{}, 
		"NamingSystem" : &NamingSystem{}, 
		"SubstanceReferenceInformationTarget" : &SubstanceReferenceInformationTarget{}, 
		"MessageHeader" : &MessageHeader{}, 
		"RiskEvidenceSynthesisSampleSize" : &RiskEvidenceSynthesisSampleSize{}, 
		"RiskEvidenceSynthesisRiskEstimate" : &RiskEvidenceSynthesisRiskEstimate{}, 
		"SubstancePolymerMonomerSet" : &SubstancePolymerMonomerSet{}, 
		"CoverageEligibilityResponse" : &CoverageEligibilityResponse{}, 
		"ValueSetCompose" : &ValueSetCompose{}, 
		"RelatedArtifact" : &RelatedArtifact{}, 
		"ConsentVerification" : &ConsentVerification{}, 
		"ContractFriendly" : &ContractFriendly{}, 
		"CoverageEligibilityRequestDiagnosis" : &CoverageEligibilityRequestDiagnosis{}, 
		"Parameters" : &Parameters{}, 
		"SubstanceNucleicAcid" : &SubstanceNucleicAcid{}, 
		"Population" : &Population{}, 
		"ClaimResponseInsurance" : &ClaimResponseInsurance{}, 
		"EffectEvidenceSynthesisResultsByExposure" : &EffectEvidenceSynthesisResultsByExposure{}, 
		"ExampleScenarioOperation" : &ExampleScenarioOperation{}, 
		"PlanDefinitionParticipant" : &PlanDefinitionParticipant{}, 
		"RiskEvidenceSynthesisCertaintySubcomponent" : &RiskEvidenceSynthesisCertaintySubcomponent{}, 
		"CodeSystem" : &CodeSystem{}, 
		"EffectEvidenceSynthesisSampleSize" : &EffectEvidenceSynthesisSampleSize{}, 
		"MedicationKnowledgeMonograph" : &MedicationKnowledgeMonograph{}, 
		"MedicinalProductAuthorization" : &MedicinalProductAuthorization{}, 
		"InsurancePlanBenefit" : &InsurancePlanBenefit{}, 
		"ObservationDefinitionQuantitativeDetails" : &ObservationDefinitionQuantitativeDetails{}, 
		"ResearchStudyObjective" : &ResearchStudyObjective{}, 
		"QuestionnaireItem" : &QuestionnaireItem{}, 
		"DataRequirementDateFilter" : &DataRequirementDateFilter{}, 
		"ListEntry" : &ListEntry{}, 
		"VerificationResult" : &VerificationResult{}, 
		"HealthcareServiceEligibility" : &HealthcareServiceEligibility{}, 
		"MolecularSequence" : &MolecularSequence{}, 
		"CoverageEligibilityResponseBenefit" : &CoverageEligibilityResponseBenefit{}, 
		"DeviceDefinitionSpecialization" : &DeviceDefinitionSpecialization{}, 
		"DeviceUseStatement" : &DeviceUseStatement{}, 
		"EvidenceVariableCharacteristic" : &EvidenceVariableCharacteristic{}, 
		"BodyStructure" : &BodyStructure{}, 
		"MedicinalProductPharmaceuticalTargetSpecies" : &MedicinalProductPharmaceuticalTargetSpecies{}, 
		"RelatedPerson" : &RelatedPerson{}, 
		"SubstanceReferenceInformation" : &SubstanceReferenceInformation{}, 
		"ClaimRelated" : &ClaimRelated{}, 
		"Device" : &Device{}, 
		"EpisodeOfCareStatusHistory" : &EpisodeOfCareStatusHistory{}, 
		"PaymentReconciliation" : &PaymentReconciliation{}, 
		"SubstanceNucleicAcidSubunit" : &SubstanceNucleicAcidSubunit{}, 
		"UsageContext" : &UsageContext{}, 
		"BiologicallyDerivedProductProcessing" : &BiologicallyDerivedProductProcessing{}, 
		"RelatedPersonCommunication" : &RelatedPersonCommunication{}, 
		"ContractTerm" : &ContractTerm{}, 
		"ContractContext" : &ContractContext{}, 
		"RiskAssessment" : &RiskAssessment{}, 
		"ClaimProcedure" : &ClaimProcedure{}, 
		"RiskEvidenceSynthesisPrecisionEstimate" : &RiskEvidenceSynthesisPrecisionEstimate{}, 
		"CapabilityStatementInteraction1" : &CapabilityStatementInteraction1{}, 
		"CatalogEntry" : &CatalogEntry{}, 
		"ClaimResponseSubDetail1" : &ClaimResponseSubDetail1{}, 
		"CodeSystemProperty1" : &CodeSystemProperty1{}, 
		"Attachment" : &Attachment{}, 
		"AllergyIntolerance" : &AllergyIntolerance{}, 
		"Binary" : &Binary{}, 
		"ImplementationGuideGlobal" : &ImplementationGuideGlobal{}, 
		"MeasureComponent" : &MeasureComponent{}, 
		"MedicinalProductInteraction" : &MedicinalProductInteraction{}, 
		"SubstanceSpecificationStructure" : &SubstanceSpecificationStructure{}, 
		"SubstanceSpecificationCode" : &SubstanceSpecificationCode{}, 
		"BiologicallyDerivedProduct" : &BiologicallyDerivedProduct{}, 
		"MessageDefinitionFocus" : &MessageDefinitionFocus{}, 
		"StructureMapTarget" : &StructureMapTarget{}, 
		"SubstancePolymerDegreeOfPolymerisation" : &SubstancePolymerDegreeOfPolymerisation{}, 
		"ImplementationGuideParameter" : &ImplementationGuideParameter{}, 
		"QuestionnaireResponseItem" : &QuestionnaireResponseItem{}, 
		"SubstanceSpecificationIsotope" : &SubstanceSpecificationIsotope{}, 
		"AllergyIntoleranceReaction" : &AllergyIntoleranceReaction{}, 
		"ClaimResponseDetail1" : &ClaimResponseDetail1{}, 
		"EvidenceVariable" : &EvidenceVariable{}, 
		"OperationDefinitionReferencedFrom" : &OperationDefinitionReferencedFrom{}, 
		"StructureMapParameter" : &StructureMapParameter{}, 
		"TestScriptAction1" : &TestScriptAction1{}, 
		"AuditEventEntity" : &AuditEventEntity{}, 
		"DetectedIssueEvidence" : &DetectedIssueEvidence{}, 
		"EnrollmentResponse" : &EnrollmentResponse{}, 
		"MedicinalProductPackagedPackageItem" : &MedicinalProductPackagedPackageItem{}, 
		"MedicinalProductPackaged" : &MedicinalProductPackaged{}, 
		"PatientContact" : &PatientContact{}, 
		"SupplyRequestParameter" : &SupplyRequestParameter{}, 
		"CapabilityStatementDocument" : &CapabilityStatementDocument{}, 
		"ContractAnswer" : &ContractAnswer{}, 
		"DeviceMetricCalibration" : &DeviceMetricCalibration{}, 
		"ImmunizationRecommendationRecommendation" : &ImmunizationRecommendationRecommendation{}, 
		"MedicinalProductManufacturingBusinessOperation" : &MedicinalProductManufacturingBusinessOperation{}, 
		"MolecularSequenceQuality" : &MolecularSequenceQuality{}, 
		"PatientCommunication" : &PatientCommunication{}, 
		"Contract" : &Contract{}, 
		"ContractRule" : &ContractRule{}, 
		"DeviceDefinitionMaterial" : &DeviceDefinitionMaterial{}, 
		"DiagnosticReport" : &DiagnosticReport{}, 
		"Composition" : &Composition{}, 
		"Patient" : &Patient{}, 
		"TestReportParticipant" : &TestReportParticipant{}, 
		"ValueSetContains" : &ValueSetContains{}, 
		"TerminologyCapabilitiesFilter" : &TerminologyCapabilitiesFilter{}, 
		"ElementDefinitionSlicing" : &ElementDefinitionSlicing{}, 
		"ImmunizationEvaluation" : &ImmunizationEvaluation{}, 
		"MedicationKnowledgeMonitoringProgram" : &MedicationKnowledgeMonitoringProgram{}, 
		"ResearchElementDefinition" : &ResearchElementDefinition{}, 
		"Questionnaire" : &Questionnaire{}, 
		"SubstanceSourceMaterialOrganism" : &SubstanceSourceMaterialOrganism{}, 
		"Task" : &Task{}, 
		"Coding" : &Coding{}, 
		"CoverageEligibilityRequest" : &CoverageEligibilityRequest{}, 
		"ImplementationGuideDefinition" : &ImplementationGuideDefinition{}, 
		"InsurancePlanGeneralCost" : &InsurancePlanGeneralCost{}, 
		"Ratio" : &Ratio{}, 
		"HumanName" : &HumanName{}, 
		"MedicationKnowledgeMedicineClassification" : &MedicationKnowledgeMedicineClassification{}, 
		"PersonLink" : &PersonLink{}, 
		"ExampleScenarioContainedInstance" : &ExampleScenarioContainedInstance{}, 
		"ImmunizationReaction" : &ImmunizationReaction{}, 
		"MedicationDispenseSubstitution" : &MedicationDispenseSubstitution{}, 
		"MolecularSequenceRepository" : &MolecularSequenceRepository{}, 
		"CoverageEligibilityResponseError" : &CoverageEligibilityResponseError{}, 
		"SpecimenDefinitionHandling" : &SpecimenDefinitionHandling{}, 
		"TaskOutput" : &TaskOutput{}, 
		"StructureMapDependent" : &StructureMapDependent{}, 
		"ElementDefinitionDiscriminator" : &ElementDefinitionDiscriminator{}, 
		"Appointment" : &Appointment{}, 
		"CoverageClass" : &CoverageClass{}, 
		"MedicinalProductContraindication" : &MedicinalProductContraindication{}, 
		"SearchParameterComponent" : &SearchParameterComponent{}, 
		"ProdCharacteristic" : &ProdCharacteristic{}, 
		"GuidanceResponse" : &GuidanceResponse{}, 
		"MeasureReportPopulation1" : &MeasureReportPopulation1{}, 
		"RiskEvidenceSynthesis" : &RiskEvidenceSynthesis{}, 
		"TerminologyCapabilities" : &TerminologyCapabilities{}, 
		"CodeSystemProperty" : &CodeSystemProperty{}, 
		"DiagnosticReportMedia" : &DiagnosticReportMedia{}, 
		"StructureMapRule" : &StructureMapRule{}, 
		"Observation" : &Observation{}, 
		"PractitionerRoleAvailableTime" : &PractitionerRoleAvailableTime{}, 
		"RiskAssessmentPrediction" : &RiskAssessmentPrediction{}, 
		"Address" : &Address{}, 
		"ElementDefinitionConstraint" : &ElementDefinitionConstraint{}, 
		"ClaimSupportingInfo" : &ClaimSupportingInfo{}, 
		"EffectEvidenceSynthesisPrecisionEstimate" : &EffectEvidenceSynthesisPrecisionEstimate{}, 
		"MeasureStratifier" : &MeasureStratifier{}, 
		"SubstancePolymerStartingMaterial" : &SubstancePolymerStartingMaterial{}, 
		"ContactDetail" : &ContactDetail{}, 
		"CapabilityStatementSecurity" : &CapabilityStatementSecurity{}, 
		"CareTeam" : &CareTeam{}, 
		"DeviceProperty" : &DeviceProperty{}, 
		"TestScriptAssert" : &TestScriptAssert{}, 
		"ContractAction" : &ContractAction{}, 
		"CoverageEligibilityResponseInsurance" : &CoverageEligibilityResponseInsurance{}, 
		"MedicinalProduct" : &MedicinalProduct{}, 
		"MolecularSequenceReferenceSeq" : &MolecularSequenceReferenceSeq{}, 
		"AppointmentParticipant" : &AppointmentParticipant{}, 
		"ExplanationOfBenefit" : &ExplanationOfBenefit{}, 
		"GraphDefinitionTarget" : &GraphDefinitionTarget{}, 
		"EncounterLocation" : &EncounterLocation{}, 
		"EpisodeOfCare" : &EpisodeOfCare{}, 
		"ImagingStudyPerformer" : &ImagingStudyPerformer{}, 
		"InsurancePlan" : &InsurancePlan{}, 
		"BiologicallyDerivedProductStorage" : &BiologicallyDerivedProductStorage{}, 
		"CommunicationPayload" : &CommunicationPayload{}, 
		"Consent" : &Consent{}, 
		"ContractSubject" : &ContractSubject{}, 
		"Location" : &Location{}, 
		"MedicationBatch" : &MedicationBatch{}, 
		"MedicationKnowledgePatientCharacteristics" : &MedicationKnowledgePatientCharacteristics{}, 
		"StructureDefinitionContext" : &StructureDefinitionContext{}, 
		"SubstanceSpecificationName" : &SubstanceSpecificationName{}, 
		"BundleEntry" : &BundleEntry{}, 
		"EncounterDiagnosis" : &EncounterDiagnosis{}, 
		"MedicationAdministrationDosage" : &MedicationAdministrationDosage{}, 
		"MedicationKnowledge" : &MedicationKnowledge{}, 
		"SubstanceReferenceInformationClassification" : &SubstanceReferenceInformationClassification{}, 
		"SubstanceSpecificationProperty" : &SubstanceSpecificationProperty{}, 
		"TestReportAssert" : &TestReportAssert{}, 
		"DataRequirementSort" : &DataRequirementSort{}, 
		"CapabilityStatementEndpoint" : &CapabilityStatementEndpoint{}, 
		"ClinicalImpressionFinding" : &ClinicalImpressionFinding{}, 
		"PlanDefinitionTarget" : &PlanDefinitionTarget{}, 
		"SupplyDelivery" : &SupplyDelivery{}, 
		"TestScriptCapability" : &TestScriptCapability{}, 
		"ImplementationGuideGrouping" : &ImplementationGuideGrouping{}, 
		"MedicinalProductSpecialDesignation" : &MedicinalProductSpecialDesignation{}, 
		"PlanDefinitionAction" : &PlanDefinitionAction{}, 
		"SpecimenDefinitionAdditive" : &SpecimenDefinitionAdditive{}, 
		"MeasurePopulation" : &MeasurePopulation{}, 
		"RequestGroupCondition" : &RequestGroupCondition{}, 
		"RiskEvidenceSynthesisCertainty" : &RiskEvidenceSynthesisCertainty{}, 
		"StructureMapStructure" : &StructureMapStructure{}, 
		"ContactPoint" : &ContactPoint{}, 
		"CompositionSection" : &CompositionSection{}, 
		"ImplementationGuidePage" : &ImplementationGuidePage{}, 
		"InsurancePlanCost" : &InsurancePlanCost{}, 
		"TestScript" : &TestScript{}, 
		"InsurancePlanLimit" : &InsurancePlanLimit{}, 
		"MedicinalProductIngredientReferenceStrength" : &MedicinalProductIngredientReferenceStrength{}, 
		"ConditionStage" : &ConditionStage{}, 
		"ContractValuedItem" : &ContractValuedItem{}, 
		"EncounterParticipant" : &EncounterParticipant{}, 
		"ImplementationGuideResource1" : &ImplementationGuideResource1{}, 
		"MedicationKnowledgeMaxDispense" : &MedicationKnowledgeMaxDispense{}, 
		"OperationDefinition" : &OperationDefinition{}, 
		"Substance" : &Substance{}, 
		"TerminologyCapabilitiesTranslation" : &TerminologyCapabilitiesTranslation{}, 
		"DocumentReferenceRelatesTo" : &DocumentReferenceRelatesTo{}, 
		"ExplanationOfBenefitAdjudication" : &ExplanationOfBenefitAdjudication{}, 
		"Flag" : &Flag{}, 
		"ImagingStudyInstance" : &ImagingStudyInstance{}, 
		"DocumentManifestRelated" : &DocumentManifestRelated{}, 
		"List" : &List{}, 
		"MeasureReportStratifier" : &MeasureReportStratifier{}, 
		"SpecimenCollection" : &SpecimenCollection{}, 
		"Contributor" : &Contributor{}, 
		"Communication" : &Communication{}, 
		"CoverageException" : &CoverageException{}, 
		"DeviceDefinition" : &DeviceDefinition{}, 
		"PaymentNotice" : &PaymentNotice{}, 
		"EncounterStatusHistory" : &EncounterStatusHistory{}, 
		"ExampleScenarioActor" : &ExampleScenarioActor{}, 
		"MolecularSequenceRoc" : &MolecularSequenceRoc{}, 
		"MedicinalProductIngredientSubstance" : &MedicinalProductIngredientSubstance{}, 
		"Organization" : &Organization{}, 
		"VerificationResultAttestation" : &VerificationResultAttestation{}, 
		"AuditEventAgent" : &AuditEventAgent{}, 
		"DetectedIssue" : &DetectedIssue{}, 
		"ImplementationGuideTemplate" : &ImplementationGuideTemplate{}, 
		"MedicinalProductIngredientSpecifiedSubstance" : &MedicinalProductIngredientSpecifiedSubstance{}, 
		"CompartmentDefinitionResource" : &CompartmentDefinitionResource{}, 
		"SubstanceSourceMaterialHybrid" : &SubstanceSourceMaterialHybrid{}, 
		"TerminologyCapabilitiesSoftware" : &TerminologyCapabilitiesSoftware{}, 
		"Distance" : &Distance{}, 
		"ClaimResponseItem" : &ClaimResponseItem{}, 
		"DocumentManifest" : &DocumentManifest{}, 
		"MeasureSupplementalData" : &MeasureSupplementalData{}, 
		"ValueSetFilter" : &ValueSetFilter{}, 
		"NutritionOrderTexture" : &NutritionOrderTexture{}, 
		"AuditEventNetwork" : &AuditEventNetwork{}, 
		"ConceptMapTarget" : &ConceptMapTarget{}, 
		"ExampleScenarioStep" : &ExampleScenarioStep{}, 
		"ImplementationGuideManifest" : &ImplementationGuideManifest{}, 
		"DocumentReferenceContent" : &DocumentReferenceContent{}, 
		"ExampleScenarioProcess" : &ExampleScenarioProcess{}, 
		"OperationOutcomeIssue" : &OperationOutcomeIssue{}, 
		"ProvenanceEntity" : &ProvenanceEntity{}, 
		"TestScriptLink" : &TestScriptLink{}, 
		"Basic" : &Basic{}, 
		"ContractSecurityLabel" : &ContractSecurityLabel{}, 
		"EffectEvidenceSynthesisCertaintySubcomponent" : &EffectEvidenceSynthesisCertaintySubcomponent{}, 
		"GroupMember" : &GroupMember{}, 
		"TestScriptMetadata" : &TestScriptMetadata{}, 
		"ElementDefinitionMapping" : &ElementDefinitionMapping{}, 
		"DeviceMetric" : &DeviceMetric{}, 
		"MeasureReportStratum" : &MeasureReportStratum{}, 
		"MedicationKnowledgeRegulatory" : &MedicationKnowledgeRegulatory{}, 
		"CompositionEvent" : &CompositionEvent{}, 
		"EventDefinition" : &EventDefinition{}, 
		"PlanDefinition" : &PlanDefinition{}, 
		"SubstanceSourceMaterialFractionDescription" : &SubstanceSourceMaterialFractionDescription{}, 
		"TerminologyCapabilitiesCodeSystem" : &TerminologyCapabilitiesCodeSystem{}, 
		"TerminologyCapabilitiesValidateCode" : &TerminologyCapabilitiesValidateCode{}, 
		"Bundle" : &Bundle{}, 
		"Media" : &Media{}, 
		"MedicinalProductUndesirableEffect" : &MedicinalProductUndesirableEffect{}, 
		"QuestionnaireResponse" : &QuestionnaireResponse{}, 
		"AccountCoverage" : &AccountCoverage{}, 
		"MedicationStatement" : &MedicationStatement{}, 
		"FamilyMemberHistoryCondition" : &FamilyMemberHistoryCondition{}, 
		"MedicinalProductAuthorizationProcedure" : &MedicinalProductAuthorizationProcedure{}, 
		"TestScriptRequestHeader" : &TestScriptRequestHeader{}, 
		"Annotation" : &Annotation{}, 
		"TriggerDefinition" : &TriggerDefinition{}, 
		"BundleRequest" : &BundleRequest{}, 
		"CoverageEligibilityRequestItem" : &CoverageEligibilityRequestItem{}, 
		"VerificationResultPrimarySource" : &VerificationResultPrimarySource{}, 
		"PractitionerRole" : &PractitionerRole{}, 
		"SpecimenProcessing" : &SpecimenProcessing{}, 
		"StructureMapGroup" : &StructureMapGroup{}, 
		"TestReportAction" : &TestReportAction{}, 
		"ObservationDefinition" : &ObservationDefinition{}, 
		"PaymentReconciliationProcessNote" : &PaymentReconciliationProcessNote{}, 
		"StructureMapSource" : &StructureMapSource{}, 
		"SubscriptionChannel" : &SubscriptionChannel{}, 
		"ConsentData" : &ConsentData{}, 
		"SubstanceProteinSubunit" : &SubstanceProteinSubunit{}, 
		"ParametersParameter" : &ParametersParameter{}, 
		"PlanDefinitionGoal" : &PlanDefinitionGoal{}, 
		"SpecimenDefinitionContainer" : &SpecimenDefinitionContainer{}, 
		"SubstancePolymerRepeatUnit" : &SubstancePolymerRepeatUnit{}, 
		"GraphDefinitionCompartment" : &GraphDefinitionCompartment{}, 
		"ImmunizationProtocolApplied" : &ImmunizationProtocolApplied{}, 
		"MedicationKnowledgeDosage" : &MedicationKnowledgeDosage{}, 
		"NutritionOrderNutrient" : &NutritionOrderNutrient{}, 
		"ActivityDefinitionDynamicValue" : &ActivityDefinitionDynamicValue{}, 
		"Goal" : &Goal{}, 
		"MedicationKnowledgeKinetics" : &MedicationKnowledgeKinetics{}, 
		"MedicinalProductPharmaceutical" : &MedicinalProductPharmaceutical{}, 
		"ObservationReferenceRange" : &ObservationReferenceRange{}, 
		"Money" : &Money{}, 
		"ClaimResponseTotal" : &ClaimResponseTotal{}, 
		"ExplanationOfBenefitTotal" : &ExplanationOfBenefitTotal{}, 
		"GraphDefinition" : &GraphDefinition{}, 
		"DetectedIssueMitigation" : &DetectedIssueMitigation{}, 
		"NutritionOrderAdministration" : &NutritionOrderAdministration{}, 
		"TestScriptSetup" : &TestScriptSetup{}, 
		"Timing" : &Timing{}, 
		"CapabilityStatementSupportedMessage" : &CapabilityStatementSupportedMessage{}, 
		"ClinicalImpressionInvestigation" : &ClinicalImpressionInvestigation{}, 
		"ConceptMapElement" : &ConceptMapElement{}, 
		"ExplanationOfBenefitRelated" : &ExplanationOfBenefitRelated{}, 
		"ExplanationOfBenefitPayment" : &ExplanationOfBenefitPayment{}, 
		"ImplementationGuidePage1" : &ImplementationGuidePage1{}, 
		"SubstanceNucleicAcidLinkage" : &SubstanceNucleicAcidLinkage{}, 
		"Reference" : &Reference{}, 
		"AuditEventDetail" : &AuditEventDetail{}, 
		"CatalogEntryRelatedEntry" : &CatalogEntryRelatedEntry{}, 
		"ClaimPayee" : &ClaimPayee{}, 
		"TerminologyCapabilitiesClosure" : &TerminologyCapabilitiesClosure{}, 
		"DataRequirement" : &DataRequirement{}, 
		"ConsentPolicy" : &ConsentPolicy{}, 
		"ExplanationOfBenefitItem" : &ExplanationOfBenefitItem{}, 
		"MedicinalProductContraindicationOtherTherapy" : &MedicinalProductContraindicationOtherTherapy{}, 
		"AuditEventSource" : &AuditEventSource{}, 
		"ChargeItemDefinition" : &ChargeItemDefinition{}, 
		"ClaimItem" : &ClaimItem{}, 
		"CoverageEligibilityRequestSupportingInfo" : &CoverageEligibilityRequestSupportingInfo{}, 
		"OperationDefinitionOverload" : &OperationDefinitionOverload{}, 
		"RequestGroup" : &RequestGroup{}, 
		"StructureMap" : &StructureMap{}, 
		"StructureMapInput" : &StructureMapInput{}, 
		"AppointmentResponse" : &AppointmentResponse{}, 
		"DeviceRequest" : &DeviceRequest{}, 
		"EffectEvidenceSynthesisCertainty" : &EffectEvidenceSynthesisCertainty{}, 
		"TestScriptOperation" : &TestScriptOperation{}, 
		"ElementDefinitionType" : &ElementDefinitionType{}, 
		"DeviceVersion" : &DeviceVersion{}, 
		"ValueSetExpansion" : &ValueSetExpansion{}, 
		"FamilyMemberHistory" : &FamilyMemberHistory{}, 
		"ImmunizationPerformer" : &ImmunizationPerformer{}, 
		"QuestionnaireInitial" : &QuestionnaireInitial{}, 
		"SubstanceAmount" : &SubstanceAmount{}, 
		"CommunicationRequestPayload" : &CommunicationRequestPayload{}, 
		"CompartmentDefinition" : &CompartmentDefinition{}, 
		"Narrative" : &Narrative{}, 
		"ClaimDiagnosis" : &ClaimDiagnosis{}, 
		"MedicationDispense" : &MedicationDispense{}, 
		"CommunicationRequest" : &CommunicationRequest{}, 
		"TestReportAction2" : &TestReportAction2{}, 
		"SubstanceSpecificationMoiety" : &SubstanceSpecificationMoiety{}, 
		"SubstanceSpecificationMolecularWeight" : &SubstanceSpecificationMolecularWeight{}, 
		"TerminologyCapabilitiesVersion" : &TerminologyCapabilitiesVersion{}, 
		"SampledData" : &SampledData{}, 
		"EffectEvidenceSynthesisEffectEstimate" : &EffectEvidenceSynthesisEffectEstimate{}, 
		"ExplanationOfBenefitDetail" : &ExplanationOfBenefitDetail{}, 
		"StructureDefinitionMapping" : &StructureDefinitionMapping{}, 
		"Expression" : &Expression{}, 
		"MedicationKnowledgeIngredient" : &MedicationKnowledgeIngredient{}, 
		"OperationDefinitionParameter" : &OperationDefinitionParameter{}, 
		"ImplementationGuide" : &ImplementationGuide{}, 
		"LinkageItem" : &LinkageItem{}, 
		"MedicinalProductInteractionInteractant" : &MedicinalProductInteractionInteractant{}, 
		"Period" : &Period{}, 
		"AuditEvent" : &AuditEvent{}, 
		"CapabilityStatementOperation" : &CapabilityStatementOperation{}, 
		"ExplanationOfBenefitProcessNote" : &ExplanationOfBenefitProcessNote{}, 
		"SupplyDeliverySuppliedItem" : &SupplyDeliverySuppliedItem{}, 
		"ElementDefinitionBase" : &ElementDefinitionBase{}, 
		"EffectEvidenceSynthesis" : &EffectEvidenceSynthesis{}, 
		"ExplanationOfBenefitFinancial" : &ExplanationOfBenefitFinancial{}, 
		"TaskRestriction" : &TaskRestriction{}, 
		"TimingRepeat" : &TimingRepeat{}, 
		"MessageHeaderSource" : &MessageHeaderSource{}, 
		"ResearchStudyArm" : &ResearchStudyArm{}, 
		"SubstanceSpecification" : &SubstanceSpecification{}, 
		"ResearchSubject" : &ResearchSubject{}, 
		"ValueSetInclude" : &ValueSetInclude{}, 
		"Extension" : &Extension{}, 
		"InvoiceLineItem" : &InvoiceLineItem{}, 
		"NutritionOrderOralDiet" : &NutritionOrderOralDiet{}, 
		"ObservationComponent" : &ObservationComponent{}, 
		"ConceptMapGroup" : &ConceptMapGroup{}, 
		"ResearchDefinition" : &ResearchDefinition{}, 
		"ValueSetDesignation" : &ValueSetDesignation{}, 
		"Invoice" : &Invoice{}, 
		"MedicinalProductAuthorizationJurisdictionalAuthorization" : &MedicinalProductAuthorizationJurisdictionalAuthorization{}, 
		"PractitionerQualification" : &PractitionerQualification{}, 
		"SubstancePolymerStructuralRepresentation" : &SubstancePolymerStructuralRepresentation{}, 
		"Identifier" : &Identifier{}, 
		"ElementDefinitionExample" : &ElementDefinitionExample{}, 
		"DeviceRequestParameter" : &DeviceRequestParameter{}, 
		"VisionPrescription" : &VisionPrescription{}, 
		"LocationPosition" : &LocationPosition{}, 
		"SubstanceSourceMaterialAuthor" : &SubstanceSourceMaterialAuthor{}, 
		"TestReportTeardown" : &TestReportTeardown{}, 
		"TestScriptVariable" : &TestScriptVariable{}, 
		"AdverseEventSuspectEntity" : &AdverseEventSuspectEntity{}, 
		"CodeSystemConcept" : &CodeSystemConcept{}, 
		"MeasureGroup" : &MeasureGroup{}, 
		"Specimen" : &Specimen{}, 
		"MedicationRequest" : &MedicationRequest{}, 
		"NutritionOrder" : &NutritionOrder{}, 
		"SearchParameter" : &SearchParameter{}, 
		"TerminologyCapabilitiesImplementation" : &TerminologyCapabilitiesImplementation{}, 
		"Quantity" : &Quantity{}, 
		"Meta" : &Meta{}, 
		"Evidence" : &Evidence{}, 
		"ExampleScenario" : &ExampleScenario{}, 
		"SubstanceSourceMaterial" : &SubstanceSourceMaterial{}, 
		"ChargeItemPerformer" : &ChargeItemPerformer{}, 
		"ImagingStudy" : &ImagingStudy{}, 
		"MedicinalProductIngredientStrength" : &MedicinalProductIngredientStrength{}, 
		"MedicinalProductManufactured" : &MedicinalProductManufactured{}, 
		"TestScriptTeardown" : &TestScriptTeardown{}, 
		"ChargeItem" : &ChargeItem{}, 
		"DeviceDefinitionCapability" : &DeviceDefinitionCapability{}, 
		"EncounterHospitalization" : &EncounterHospitalization{}, 
		"QuestionnaireResponseAnswer" : &QuestionnaireResponseAnswer{}, 
		"OperationOutcome" : &OperationOutcome{}, 
		"RequestGroupRelatedAction" : &RequestGroupRelatedAction{}, 
		"Schedule" : &Schedule{}, 
		"ServiceRequest" : &ServiceRequest{}, 
		"CompositionRelatesTo" : &CompositionRelatesTo{}, 
		"ContractOffer" : &ContractOffer{}, 
		"ImagingStudySeries" : &ImagingStudySeries{}, 
		"MedicationRequestSubstitution" : &MedicationRequestSubstitution{}, 
		"AccountGuarantor" : &AccountGuarantor{}, 
		"CapabilityStatement" : &CapabilityStatement{}, 
		"InsurancePlanCoverage" : &InsurancePlanCoverage{}, 
		"Medication" : &Medication{}, 
		"MedicationKnowledgeAdministrationGuidelines" : &MedicationKnowledgeAdministrationGuidelines{}, 
		"MedicinalProductNamePart" : &MedicinalProductNamePart{}, 
		"ResearchStudy" : &ResearchStudy{}, 
		"ClaimResponse" : &ClaimResponse{}, 
		"ConceptMapUnmapped" : &ConceptMapUnmapped{}, 
		"MedicationIngredient" : &MedicationIngredient{}, 
		"MedicationKnowledgeCost" : &MedicationKnowledgeCost{}, 
		"ConceptMapDependsOn" : &ConceptMapDependsOn{}, 
		"DeviceDeviceName" : &DeviceDeviceName{}, 
		"Measure" : &Measure{}, 
		"PractitionerRoleNotAvailable" : &PractitionerRoleNotAvailable{}, 
		"QuestionnaireAnswerOption" : &QuestionnaireAnswerOption{}, 
		"SubstanceReferenceInformationGene" : &SubstanceReferenceInformationGene{}, 
		"AdverseEvent" : &AdverseEvent{}, 
		"CapabilityStatementSoftware" : &CapabilityStatementSoftware{}, 
		"ClinicalImpression" : &ClinicalImpression{}, 
		"InsurancePlanPlan" : &InsurancePlanPlan{}, 
		"TestReportSetup" : &TestReportSetup{}, 
		"DocumentReference" : &DocumentReference{}, 
		"MessageHeaderResponse" : &MessageHeaderResponse{}, 
		"NutritionOrderEnteralFormula" : &NutritionOrderEnteralFormula{}, 
		"SubstancePolymer" : &SubstancePolymer{}, 
		"HealthcareServiceAvailableTime" : &HealthcareServiceAvailableTime{}, 
		"MedicationKnowledgeSchedule" : &MedicationKnowledgeSchedule{}, 
		"TestReportAction1" : &TestReportAction1{}, 
		"MedicinalProductIndication" : &MedicinalProductIndication{}, 
		"MedicinalProductIngredient" : &MedicinalProductIngredient{}, 
		"SubstanceSpecificationRepresentation" : &SubstanceSpecificationRepresentation{}, 
		"CapabilityStatementInteraction" : &CapabilityStatementInteraction{}, 
		"DeviceUdiCarrier" : &DeviceUdiCarrier{}, 
		"EpisodeOfCareDiagnosis" : &EpisodeOfCareDiagnosis{}, 
		"ExplanationOfBenefitCareTeam" : &ExplanationOfBenefitCareTeam{}, 
		"ClaimResponsePayment" : &ClaimResponsePayment{}, 
		"MedicinalProductName" : &MedicinalProductName{}, 
		"SubstanceInstance" : &SubstanceInstance{}, 
		"MessageDefinition" : &MessageDefinition{}, 
		"Count" : &Count{}, 
		"BundleLink" : &BundleLink{}, 
		"Encounter" : &Encounter{}, 
		"EncounterClassHistory" : &EncounterClassHistory{}, 
		"TestReportOperation" : &TestReportOperation{}, 
		"TestScriptOrigin" : &TestScriptOrigin{}, 
		"DeviceDefinitionProperty" : &DeviceDefinitionProperty{}, 
		"InsurancePlanBenefit1" : &InsurancePlanBenefit1{}, 
		"StructureDefinition" : &StructureDefinition{}, 
		"TaskInput" : &TaskInput{}, 
		"SubstanceSourceMaterialPartDescription" : &SubstanceSourceMaterialPartDescription{}, 
		"ContractAsset" : &ContractAsset{}, 
		"ExplanationOfBenefitProcedure" : &ExplanationOfBenefitProcedure{}, 
		"OperationDefinitionBinding" : &OperationDefinitionBinding{}, 
		"ProvenanceAgent" : &ProvenanceAgent{}, 
		"TestScriptFixture" : &TestScriptFixture{}, 
		"DeviceSpecialization" : &DeviceSpecialization{}, 
		"MedicationKnowledgeSubstitution" : &MedicationKnowledgeSubstitution{}, 
		"MedicinalProductPackagedBatchIdentifier" : &MedicinalProductPackagedBatchIdentifier{}, 
		"TestScriptAction" : &TestScriptAction{}, 
		"ExplanationOfBenefitDetail1" : &ExplanationOfBenefitDetail1{}, 
		"InvoicePriceComponent" : &InvoicePriceComponent{}, 
		"QuestionnaireEnableWhen" : &QuestionnaireEnableWhen{}, 
		"Slot" : &Slot{}, 
		"ObservationDefinitionQualifiedInterval" : &ObservationDefinitionQualifiedInterval{}, 
		"SubstanceProtein" : &SubstanceProtein{}, 
		"TestReport" : &TestReport{}, 
		"Claim" : &Claim{}, 
		"ClaimResponseError" : &ClaimResponseError{}, 
		"ExampleScenarioInstance" : &ExampleScenarioInstance{}, 
		"ImmunizationRecommendation" : &ImmunizationRecommendation{}, 
		"ConsentProvision" : &ConsentProvision{}, 
		"ContractSigner" : &ContractSigner{}, 
		"MedicationKnowledgeDrugCharacteristic" : &MedicationKnowledgeDrugCharacteristic{}, 
		"MedicinalProductIndicationOtherTherapy" : &MedicinalProductIndicationOtherTherapy{}, 
		"DataRequirementCodeFilter" : &DataRequirementCodeFilter{}, 
		"DosageDoseAndRate" : &DosageDoseAndRate{}, 
		"CapabilityStatementRest" : &CapabilityStatementRest{}, 
		"CodeSystemDesignation" : &CodeSystemDesignation{}, 
		"TestScriptTest" : &TestScriptTest{}, 
		"MedicinalProductCountryLanguage" : &MedicinalProductCountryLanguage{}, 
		"MolecularSequenceVariant" : &MolecularSequenceVariant{}, 
		"Coverage" : &Coverage{}, 
		"ImplementationGuideResource" : &ImplementationGuideResource{}, 
		"MedicationRequestInitialFill" : &MedicationRequestInitialFill{}, 
		"DocumentReferenceContext" : &DocumentReferenceContext{}, 
		"InsurancePlanContact" : &InsurancePlanContact{}, 
		"TestScriptDestination" : &TestScriptDestination{}, 
		"ClaimAccident" : &ClaimAccident{}, 
		"ClaimSubDetail" : &ClaimSubDetail{}, 
		"ContractContentDefinition" : &ContractContentDefinition{}, 
		"CoverageCostToBeneficiary" : &CoverageCostToBeneficiary{}, 
		"Signature" : &Signature{}, 
		"MedicationKnowledgePackaging" : &MedicationKnowledgePackaging{}, 
		"Subscription" : &Subscription{}, 
		"ValueSet" : &ValueSet{}, 
		"BundleResponse" : &BundleResponse{}, 
		"DeviceDefinitionDeviceName" : &DeviceDefinitionDeviceName{}, 
		"PaymentReconciliationDetail" : &PaymentReconciliationDetail{}, 
		"StructureDefinitionSnapshot" : &StructureDefinitionSnapshot{}, 
		"ExplanationOfBenefitSupportingInfo" : &ExplanationOfBenefitSupportingInfo{}, 
		"InvoiceParticipant" : &InvoiceParticipant{}, 
		"MedicinalProductPharmaceuticalRouteOfAdministration" : &MedicinalProductPharmaceuticalRouteOfAdministration{}, 
		"PlanDefinitionCondition" : &PlanDefinitionCondition{}, 
		"BundleSearch" : &BundleSearch{}, 
		"ContractParty" : &ContractParty{}, 
		"Group" : &Group{}, 
		"MeasureReport" : &MeasureReport{}, 
		"StructureDefinitionDifferential" : &StructureDefinitionDifferential{}, 
		"SubstancePolymerRepeat" : &SubstancePolymerRepeat{}, 
		"SubstanceReferenceInformationGeneElement" : &SubstanceReferenceInformationGeneElement{}, 
		"Range" : &Range{}, 
		"SubstanceAmountReferenceRange" : &SubstanceAmountReferenceRange{}, 
		"ClaimResponseSubDetail" : &ClaimResponseSubDetail{}, 
		"EnrollmentRequest" : &EnrollmentRequest{}, 
		"SubstanceSpecificationOfficial" : &SubstanceSpecificationOfficial{}, 
		"VerificationResultValidator" : &VerificationResultValidator{}, 
		"ConsentActor" : &ConsentActor{}, 
		"GroupCharacteristic" : &GroupCharacteristic{}, 
		"HealthcareService" : &HealthcareService{}, 
		"CodeSystemFilter" : &CodeSystemFilter{}, 
		"MeasureReportGroup" : &MeasureReportGroup{}, 
		"SubstanceNucleicAcidSugar" : &SubstanceNucleicAcidSugar{}, 
		"MeasureReportComponent" : &MeasureReportComponent{}, 
		"MolecularSequenceOuter" : &MolecularSequenceOuter{}, 
		"Procedure" : &Procedure{}, 
		"SupplyRequest" : &SupplyRequest{}, 
		"Element" : &Element{}, 
		"Duration" : &Duration{}, 
		"CapabilityStatementResource" : &CapabilityStatementResource{}, 
		"ExplanationOfBenefitBenefitBalance" : &ExplanationOfBenefitBenefitBalance{}, 
		"ActivityDefinition" : &ActivityDefinition{}, 
		"Endpoint" : &Endpoint{}, 
		"SubstanceSourceMaterialOrganismGeneral" : &SubstanceSourceMaterialOrganismGeneral{}, 
		"PlanDefinitionDynamicValue" : &PlanDefinitionDynamicValue{}, 
		"TerminologyCapabilitiesParameter" : &TerminologyCapabilitiesParameter{}, 
		"ActivityDefinitionParticipant" : &ActivityDefinitionParticipant{}, 
		"ImmunizationEducation" : &ImmunizationEducation{}, 
		"MedicinalProductPharmaceuticalWithdrawalPeriod" : &MedicinalProductPharmaceuticalWithdrawalPeriod{}, 
		"PlanDefinitionRelatedAction" : &PlanDefinitionRelatedAction{}, 
		"ResearchElementDefinitionCharacteristic" : &ResearchElementDefinitionCharacteristic{}, 
		"BiologicallyDerivedProductCollection" : &BiologicallyDerivedProductCollection{}, 
		"ChargeItemDefinitionApplicability" : &ChargeItemDefinitionApplicability{}, 
		"ChargeItemDefinitionPropertyGroup" : &ChargeItemDefinitionPropertyGroup{}, 
		"ProcedurePerformer" : &ProcedurePerformer{}, 
		"OrganizationContact" : &OrganizationContact{}, 
		"Provenance" : &Provenance{}, 
		"Account" : &Account{}, 
		"CapabilityStatementMessaging" : &CapabilityStatementMessaging{}, 
		"ClaimResponseDetail" : &ClaimResponseDetail{}, 
		"ExplanationOfBenefitAccident" : &ExplanationOfBenefitAccident{}, 
		"ExplanationOfBenefitAddItem" : &ExplanationOfBenefitAddItem{}, 
		"MedicationAdministration" : &MedicationAdministration{}, 
		"MessageDefinitionAllowedResponse" : &MessageDefinitionAllowedResponse{}, 
		"OrganizationAffiliation" : &OrganizationAffiliation{}, 
		"MarketingStatus" : &MarketingStatus{}, 
		"ClaimResponseAdjudication" : &ClaimResponseAdjudication{}, 
		"CoverageEligibilityResponseItem" : &CoverageEligibilityResponseItem{}, 
		"ExampleScenarioVersion" : &ExampleScenarioVersion{}, 
		"ProcedureFocalDevice" : &ProcedureFocalDevice{}, 
		"RequestGroupAction" : &RequestGroupAction{}, 
		"CapabilityStatementImplementation" : &CapabilityStatementImplementation{}, 
		"CapabilityStatementSearchParam" : &CapabilityStatementSearchParam{}, 
		"MedicationKnowledgeRelatedMedicationKnowledge" : &MedicationKnowledgeRelatedMedicationKnowledge{}, 
		"NutritionOrderSupplement" : &NutritionOrderSupplement{}, 
	}
	resource, ok := FhirStructMap[resource_name]
	if ok == true {
		return resource, nil
	}
	return resource, errors.New("Resource not found")
}
