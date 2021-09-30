package utility_test

import (
	"fmt"
	"testing"
	"tweet-image-downloader/entity"
	"tweet-image-downloader/utility"
)

func Test_ValidateEmptyParamValue(t *testing.T) {
	tests := []struct {
		name  string
		input *utility.ParamBuilder
		want  error
	}{
		{
			name:  "failed cause of expansion Parameter Value is empty",
			input: utility.NewParamBuilder().Expansions(entity.ExpansionFields{}),
			want:  entity.ErrorEmptyParameterValue,
		},
		{
			name:  "failed cause of media.fields Parameter Value is empty",
			input: utility.NewParamBuilder().MediaFields(entity.MediaFields{}),
			want:  entity.ErrorEmptyParameterValue,
		},
		{
			name:  "failed cause of query Parameter Value is empty",
			input: utility.NewParamBuilder().Query(entity.QueryFields{}),
			want:  entity.ErrorEmptyParameterValue,
		},
		{
			name:  "failed cause of tweet.fields Parameter Value is empty",
			input: utility.NewParamBuilder().TweetFields(entity.TweetFields{}),
			want:  entity.ErrorEmptyParameterValue,
		},
		{
			name: "failed cause of tweet.fields Parameter Value is empty, but query Parameter Value is not empty",
			input: utility.NewParamBuilder().TweetFields(entity.TweetFields{}).Query(entity.QueryFields{
				entity.QueryFieldHasImages,
			}),
			want: entity.ErrorEmptyParameterValue,
		},
		{
			name:  "success only one parameter",
			input: utility.NewParamBuilder().Query(entity.QueryFields{entity.QueryFieldHasImages}),
			want:  nil,
		},
		{
			name: "success when multiple parameters",
			input: utility.NewParamBuilder().
				Query(entity.QueryFields{
					entity.QueryFieldHasImages,
					entity.QueryFieldFrom("test"),
				}).
				TweetFields(entity.TweetFields{
					entity.TweetFieldCreatedAt,
					entity.TweetFieldAttachments,
				}),
			want: nil,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, test.name), func(t *testing.T) {
			if str, err := utility.ExportValidateEmptyParamValue(test.input); err != test.want {
				t.Errorf("failed to validate: str is %v, result is %v, want is %v", str, err, test.want)
			}
		})
	}
}
