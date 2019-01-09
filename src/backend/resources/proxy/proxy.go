package proxy

import (
	"k8s.io/apimachinery/pkg/util/json"

	"github.com/Qihoo360/wayne/src/backend/client"
	"github.com/Qihoo360/wayne/src/backend/common"
	"github.com/Qihoo360/wayne/src/backend/models/response"
	"github.com/Qihoo360/wayne/src/backend/resources/dataselector"
)

func GetPage(kubeClient client.ResourceHandler, kind string, namespace string, q *common.QueryParam) (*common.Page, error) {
	objs, err := kubeClient.List(kind, namespace, q.LabelSelector)
	if err != nil {
		return nil, err
	}
	commonObjs := make([]dataselector.DataCell, 0)
	for _, obj := range objs {
		objByte, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		var commonObj ObjectCell
		err = json.Unmarshal(objByte, &commonObj)
		if err != nil {
			return nil, err
		}
		commonObjs = append(commonObjs, commonObj)
	}

	return dataselector.DataSelectPage(commonObjs, q), nil
}

func GetNames(kubeClient client.ResourceHandler, kind string, namespace string) ([]response.NamesObject, error) {
	objs, err := kubeClient.List(kind, namespace, "")
	if err != nil {
		return nil, err
	}
	commonObjs := make([]response.NamesObject, 0)
	for _, obj := range objs {
		objByte, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		var commonObj ObjectCell
		err = json.Unmarshal(objByte, &commonObj)
		if err != nil {
			return nil, err
		}
		commonObjs = append(commonObjs, response.NamesObject{
			Name: commonObj.Name,
		})
	}

	return commonObjs, nil
}