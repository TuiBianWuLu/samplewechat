package material

import (
	"fmt"
	"encoding/json"
	"github.com/TuiBianWuLu/samplewechat/core/request"
	"github.com/TuiBianWuLu/samplewechat/basic"
)

const (
	addNewsURL     = "https://api.weixin.qq.com/cgi-bin/material/add_news"
	addMaterialURL = "https://api.weixin.qq.com/cgi-bin/material/add_material"
	delMaterialURL = "https://api.weixin.qq.com/cgi-bin/material/del_material"
)

//Material 素材管理
type Material struct {
	*basic.WeChat
}

//reqArticles 永久性图文素材请求信息
type reqArticles struct {
	Articles []*Article `json:"articles"`
}
//reqDeleteMaterial 删除微信素材
type reqDeleteMaterial struct {
	MediaID string `json:"media_id"`
}

//resArticles 永久性图文素材返回结果
type resArticles struct {
	basic.CommonError
	MediaID 	string `json:"media_id"`
}

//article 新增永久素材
type Article struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	ShowCoverPic     int    `json:"show_cover_pic"`
	Content          string `json:"content"`
	ContentSourceURL string `json:"content_source_url"`
}

func (m *Material) AddNews (article []*Article) (mediaId string, err error)  {
	rep := &reqArticles{article}

	res, err := m.RefreshToken()

	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", addNewsURL, res.AccessToken)

	responseBytes, err := request.Post(uri, rep)

	var result resArticles

	err = json.Unmarshal(responseBytes, result)
	if err != nil {
		return
	}
	mediaId = result.MediaID
	return

}

func (m *Material) Delete (mediaId string) error {
	res, err := m.RefreshToken()

	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", delMaterialURL, res.AccessToken)

	rep := &reqDeleteMaterial{mediaId}

	body, err := request.Post(uri, rep)

	if err != nil {
		return err
	}

	var result basic.CommonError

	err = json.Unmarshal(body, result)

	if err != nil {
		return err
	}

	if result.ErrCode != '0' {
		return fmt.Errorf("delete materila fail error msg : %s", result.ErrMsg)
	}
	return nil
}

func (m *Material) AddMaterial (mType MediaType, filename string) error {

	res, err := m.RefreshToken()

	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", addMaterialURL, res.AccessToken)

	return nil
}