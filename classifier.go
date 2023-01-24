package catboostcpai

// Classifier is wrapper over model object that add methods for binary classification
type Classifier struct {
	Model *Model
}

// LoadClassifierFromFile loads binary classifier from file
func LoadClassifierFromFile(filename string) (*Classifier, error) {
	model, err := LoadFullModelFromFile(filename)
	if err != nil {
		return nil, err
	}
	return &Classifier{Model: model}, nil
}

// Predict returns scores
func (bc *Classifier) Predict(floatFeatures [][]float32, floatFeaturesLength int, catFeatures [][]string, catFeaturesLength int) ([]float64, error) {
	results, err := bc.Model.CalcModelPrediction(floatFeatures, floatFeaturesLength, catFeatures, catFeaturesLength)
	if err != nil {
		return nil, err
	}
	return results, nil
}
