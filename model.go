package catboostcapi

/*
#cgo linux LDFLAGS: -lcatboostmodel
#cgo darwin LDFLAGS: -lcatboostmodel
#include <stdlib.h>
#include <stdbool.h>
#include <model_calcer_wrapper.h>

static char** makeCharArray(int size) {
    return calloc(sizeof(char*), size);
}

static void freeCharArray(char **a, int size) {
	int i;
	for (i = 0; i < size; i++)
		free(a[i]);
	free(a);
}

static char*** makeArrayCharArray(int size) {
	return calloc(sizeof(char**), size);
}

static void freeArrayCharArray(char ***a, int sizeX, int sizeY) {
	int i;
	for (i = 0; i < sizeX; i++)
		freeCharArray(a[i], sizeY);
	free(a);
}

static float** makeFloatArray(int size) {
	return calloc(sizeof(float*), size);
}

static void setArrayCharArray(char ***a, char **s, int n){
	a[n] = s;
}

static void setFloatArray(float **a, float *f, int n){
	a[n] = f;
}

static void setCharArray(char **a, char *s, int n) {
    a[n] = s;
}

*/
import "C"

import (
	"fmt"
	"unsafe"
)

func getError() error {
	messageC := C.GetErrorString()
	message := C.GoString(messageC)
	return fmt.Errorf(message)
}

// Model is a wrapper over ModelCalcerHandler
type Model struct {
	Handler unsafe.Pointer
}

// GetFloatFeaturesCount returns a number of float features used for training
func (model *Model) GetFloatFeaturesCount() int {
	return int(C.GetFloatFeaturesCount(model.Handler))
}

// GetCatFeaturesCount returns a number of categorical features used for training
func (model *Model) GetCatFeaturesCount() int {
	return int(C.GetCatFeaturesCount(model.Handler))
}

// Close deletes model handler
func (model *Model) Close() {
	C.ModelCalcerDelete(model.Handler)
}

// LoadFullModelFromFile loads model from file
func LoadFullModelFromFile(filename string) (*Model, error) {
	model := &Model{}
	model.Handler = C.ModelCalcerCreate()
	if !C.LoadFullModelFromFile(model.Handler, C.CString(filename)) {
		return nil, getError()
	}
	return model, nil
}

// CalcModelPrediction returns raw predictions for specified data points
func (model *Model) CalcModelPrediction(floatFeatures [][]float32, floatFeaturesLength int, catFeatures [][]string, catFeaturesLength int) ([]float64, error) {
	nSamples := len(floatFeatures)
	results := make([]float64, nSamples)

	// Allocates memory for an array of num objects of size
	// and initializes all bytes in the allocated storage to zero
	floatsC := C.makeFloatArray(C.int(nSamples))

	// dynamically de-allocate the memory
	defer C.free(unsafe.Pointer(floatsC))

	// set go ptr to c ptr
	for i, v := range floatFeatures {
		C.setFloatArray(floatsC, (*C.float)(&v[0]), C.int(i))
	}

	// Allocates memory for an array of num objects of size
	// and initializes all bytes in the allocated storage to zero
	catsC := C.makeArrayCharArray(C.int(nSamples))

	// dynamically de-allocate the memory
	defer C.freeArrayCharArray(catsC, C.int(nSamples), C.int(catFeaturesLength))

	// generate C array of array of chars
	for i, cat := range catFeatures {
		catC := C.makeCharArray(C.int(len(cat)))
		for i, c := range cat {
			C.setCharArray(catC, C.CString(c), C.int(i))
		}
		C.setArrayCharArray(catsC, catC, C.int(i))
	}

	if !C.CalcModelPrediction(
		model.Handler,
		C.size_t(nSamples),
		floatsC,
		C.size_t(floatFeaturesLength),
		catsC,
		C.size_t(catFeaturesLength),
		(*C.double)(&results[0]),
		C.size_t(nSamples),
	) {
		return nil, getError()
	}

	return results, nil
}
