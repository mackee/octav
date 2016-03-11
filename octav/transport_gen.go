// Automatically generated by gentransport utility. DO NOT EDIT!
package octav

import (
	"encoding/json"
	"errors"

	"github.com/lestrrat/go-urlenc"
)

func (r CreateSessionRequest) collectMarshalData() map[string]interface{} {
	m := make(map[string]interface{})
	m["conference_id"] = r.ConferenceID
	m["speaker_id"] = r.SpeakerID
	m["title"] = r.Title
	if r.Abstract.Valid() {
		m["abstract"] = r.Abstract.Value()
	}
	if r.Memo.Valid() {
		m["memo"] = r.Memo.Value()
	}
	m["duration"] = r.Duration
	if r.MaterialLevel.Valid() {
		m["material_level"] = r.MaterialLevel.Value()
	}
	m["tags"] = r.Tags
	if r.Category.Valid() {
		m["category"] = r.Category.Value()
	}
	if r.SpokenLanguage.Valid() {
		m["spoken_language"] = r.SpokenLanguage.Value()
	}
	if r.SlideLanguage.Valid() {
		m["slide_language"] = r.SlideLanguage.Value()
	}
	if r.SlideSubtitles.Valid() {
		m["slide_subtitles"] = r.SlideSubtitles.Value()
	}
	if r.SlideURL.Valid() {
		m["slide_url"] = r.SlideURL.Value()
	}
	if r.VideoURL.Valid() {
		m["video_url"] = r.VideoURL.Value()
	}
	if r.PhotoPermission.Valid() {
		m["photo_permission"] = r.PhotoPermission.Value()
	}
	if r.VideoPermission.Valid() {
		m["video_permission"] = r.VideoPermission.Value()
	}
	return m
}

func (r CreateSessionRequest) MarshalJSON() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return marshalJSONWithL10N(buf, r.L10N)
}

func (r CreateSessionRequest) MarshalURL() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := urlenc.Marshal(m)
	if err != nil {
		return nil, err
	}
	return marshalURLWithL10N(buf, r.L10N)
}

func (r *CreateSessionRequest) UnmarshalJSON(data []byte) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if jv, ok := m["conference_id"]; ok {
		switch jv.(type) {
		case string:
			r.ConferenceID = jv.(string)
			delete(m, "conference_id")
		default:
			return ErrInvalidJSONFieldType{Field: "conference_id"}
		}
	}
	if jv, ok := m["speaker_id"]; ok {
		switch jv.(type) {
		case string:
			r.SpeakerID = jv.(string)
			delete(m, "speaker_id")
		default:
			return ErrInvalidJSONFieldType{Field: "speaker_id"}
		}
	}
	if jv, ok := m["title"]; ok {
		switch jv.(type) {
		case string:
			r.Title = jv.(string)
			delete(m, "title")
		default:
			return ErrInvalidJSONFieldType{Field: "title"}
		}
	}
	if jv, ok := m["abstract"]; ok {
		r.Abstract.Set(jv)
		delete(m, "abstract")
	}
	if jv, ok := m["memo"]; ok {
		r.Memo.Set(jv)
		delete(m, "memo")
	}
	if jv, ok := m["duration"]; ok {
		switch jv.(type) {
		case int:
			r.Duration = jv.(int)
			delete(m, "duration")
		default:
			return ErrInvalidJSONFieldType{Field: "duration"}
		}
	}
	if jv, ok := m["material_level"]; ok {
		r.MaterialLevel.Set(jv)
		delete(m, "material_level")
	}
	if jv, ok := m["tags"]; ok {
		switch jv.(type) {
		case []string:
			r.Tags = jv.([]string)
			delete(m, "tags")
		default:
			return ErrInvalidJSONFieldType{Field: "tags"}
		}
	}
	if jv, ok := m["category"]; ok {
		r.Category.Set(jv)
		delete(m, "category")
	}
	if jv, ok := m["spoken_language"]; ok {
		r.SpokenLanguage.Set(jv)
		delete(m, "spoken_language")
	}
	if jv, ok := m["slide_language"]; ok {
		r.SlideLanguage.Set(jv)
		delete(m, "slide_language")
	}
	if jv, ok := m["slide_subtitles"]; ok {
		r.SlideSubtitles.Set(jv)
		delete(m, "slide_subtitles")
	}
	if jv, ok := m["slide_url"]; ok {
		r.SlideURL.Set(jv)
		delete(m, "slide_url")
	}
	if jv, ok := m["video_url"]; ok {
		r.VideoURL.Set(jv)
		delete(m, "video_url")
	}
	if jv, ok := m["photo_permission"]; ok {
		r.PhotoPermission.Set(jv)
		delete(m, "photo_permission")
	}
	if jv, ok := m["video_permission"]; ok {
		r.VideoPermission.Set(jv)
		delete(m, "video_permission")
	}
	if err := ExtractL10NFields(m, &r.L10N, []string{"conference_id", "speaker_id", "title", "abstract", "memo", "duration", "material_level", "tags", "category", "spoken_language", "slide_language", "slide_subtitles", "slide_url", "video_url", "photo_permission", "video_permission"}); err != nil {
		return err
	}
	return nil
}

func (r *CreateSessionRequest) GetPropNames() ([]string, error) {
	l, _ := r.L10N.GetPropNames()
	return append(l, "conference_id", "speaker_id", "title", "abstract", "memo", "duration", "material_level", "tags", "category", "spoken_language", "slide_language", "slide_subtitles", "slide_url", "video_url", "photo_permission", "video_permission"), nil
}

func (r *CreateSessionRequest) SetPropValue(s string, v interface{}) error {
	switch s {
	case "conference_id":
		if jv, ok := v.(string); ok {
			r.ConferenceID = jv
			return nil
		}
	case "speaker_id":
		if jv, ok := v.(string); ok {
			r.SpeakerID = jv
			return nil
		}
	case "title":
		if jv, ok := v.(string); ok {
			r.Title = jv
			return nil
		}
	case "abstract":
		return r.Abstract.Set(v)
	case "memo":
		return r.Memo.Set(v)
	case "duration":
		if jv, ok := v.(int); ok {
			r.Duration = jv
			return nil
		}
	case "material_level":
		return r.MaterialLevel.Set(v)
	case "tags":
		if jv, ok := v.([]string); ok {
			r.Tags = jv
			return nil
		}
	case "category":
		return r.Category.Set(v)
	case "spoken_language":
		return r.SpokenLanguage.Set(v)
	case "slide_language":
		return r.SlideLanguage.Set(v)
	case "slide_subtitles":
		return r.SlideSubtitles.Set(v)
	case "slide_url":
		return r.SlideURL.Set(v)
	case "video_url":
		return r.VideoURL.Set(v)
	case "photo_permission":
		return r.PhotoPermission.Set(v)
	case "video_permission":
		return r.VideoPermission.Set(v)
	default:
		return errors.New("unknown column '" + s + "'")
	}
	return ErrInvalidFieldType{Field: s}
}

func (r ListVenueRequest) collectMarshalData() map[string]interface{} {
	m := make(map[string]interface{})
	if r.Since.Valid() {
		m["since"] = r.Since.Value()
	}
	if r.Lang.Valid() {
		m["lang"] = r.Lang.Value()
	}
	if r.Limit.Valid() {
		m["limit"] = r.Limit.Value()
	}
	return m
}

func (r ListVenueRequest) MarshalJSON() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (r ListVenueRequest) MarshalURL() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := urlenc.Marshal(m)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (r *ListVenueRequest) UnmarshalJSON(data []byte) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if jv, ok := m["since"]; ok {
		r.Since.Set(jv)
		delete(m, "since")
	}
	if jv, ok := m["lang"]; ok {
		r.Lang.Set(jv)
		delete(m, "lang")
	}
	if jv, ok := m["limit"]; ok {
		r.Limit.Set(jv)
		delete(m, "limit")
	}
	return nil
}

func (r UpdateConferenceRequest) collectMarshalData() map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = r.ID
	if r.Title.Valid() {
		m["title"] = r.Title.Value()
	}
	if r.SubTitle.Valid() {
		m["sub_title"] = r.SubTitle.Value()
	}
	if r.Slug.Valid() {
		m["slug"] = r.Slug.Value()
	}
	return m
}

func (r UpdateConferenceRequest) MarshalJSON() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return marshalJSONWithL10N(buf, r.L10N)
}

func (r UpdateConferenceRequest) MarshalURL() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := urlenc.Marshal(m)
	if err != nil {
		return nil, err
	}
	return marshalURLWithL10N(buf, r.L10N)
}

func (r *UpdateConferenceRequest) UnmarshalJSON(data []byte) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if jv, ok := m["id"]; ok {
		switch jv.(type) {
		case string:
			r.ID = jv.(string)
			delete(m, "id")
		default:
			return ErrInvalidJSONFieldType{Field: "id"}
		}
	}
	if jv, ok := m["title"]; ok {
		r.Title.Set(jv)
		delete(m, "title")
	}
	if jv, ok := m["sub_title"]; ok {
		r.SubTitle.Set(jv)
		delete(m, "sub_title")
	}
	if jv, ok := m["slug"]; ok {
		r.Slug.Set(jv)
		delete(m, "slug")
	}
	if err := ExtractL10NFields(m, &r.L10N, []string{"id", "title", "sub_title", "slug"}); err != nil {
		return err
	}
	return nil
}

func (r *UpdateConferenceRequest) GetPropNames() ([]string, error) {
	l, _ := r.L10N.GetPropNames()
	return append(l, "id", "title", "sub_title", "slug"), nil
}

func (r *UpdateConferenceRequest) SetPropValue(s string, v interface{}) error {
	switch s {
	case "id":
		if jv, ok := v.(string); ok {
			r.ID = jv
			return nil
		}
	case "title":
		return r.Title.Set(v)
	case "sub_title":
		return r.SubTitle.Set(v)
	case "slug":
		return r.Slug.Set(v)
	default:
		return errors.New("unknown column '" + s + "'")
	}
	return ErrInvalidFieldType{Field: s}
}

func (r ListSessionsByConferenceRequest) collectMarshalData() map[string]interface{} {
	m := make(map[string]interface{})
	m["conference_id"] = r.ConferenceID
	if r.Date.Valid() {
		m["date"] = r.Date.Value()
	}
	return m
}

func (r ListSessionsByConferenceRequest) MarshalJSON() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (r ListSessionsByConferenceRequest) MarshalURL() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := urlenc.Marshal(m)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (r *ListSessionsByConferenceRequest) UnmarshalJSON(data []byte) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if jv, ok := m["conference_id"]; ok {
		switch jv.(type) {
		case string:
			r.ConferenceID = jv.(string)
			delete(m, "conference_id")
		default:
			return ErrInvalidJSONFieldType{Field: "conference_id"}
		}
	}
	if jv, ok := m["date"]; ok {
		r.Date.Set(jv)
		delete(m, "date")
	}
	return nil
}
