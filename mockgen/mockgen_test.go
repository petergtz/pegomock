package mockgen_test

import (
	"github.com/petergtz/pegomock/mockgen"
	"github.com/petergtz/pegomock/modelgen/loader"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mockgen", func() {
	Context("matcherSourceCodes", func() {
		It("uses correct naming pattern with underscores for keys, and correct types etc. in source code", func() {
			ast, e := loader.GenerateModel("github.com/petergtz/pegomock/test_interface", "Display")
			Expect(e).NotTo(HaveOccurred())
			_, matcherSourceCodes := mockgen.GenerateOutput(ast, "irrelevant", "MockDisplay", "test_package", "")

			Expect(matcherSourceCodes).To(SatisfyAll(
				HaveLen(12),
				HaveKeyWithValue("http_request", SatisfyAll(
					ContainSubstring("http \"net/http\""),
					ContainSubstring("func AnyHttpRequest() http.Request"),
				)),
				HaveKeyWithValue("ptr_to_http_request", SatisfyAll(
					ContainSubstring("http \"net/http\""),
					ContainSubstring("func AnyPtrToHttpRequest() *http.Request"),
				)),
				HaveKeyWithValue("slice_of_string",
					ContainSubstring("func AnySliceOfString() []string"),
				),
				HaveKeyWithValue("map_of_string_to_http_request", SatisfyAll(
					ContainSubstring("http \"net/http\""),
					ContainSubstring("func AnyMapOfStringToHttpRequest() map[string]http.Request"),
				)),
				HaveKeyWithValue("io_readcloser", SatisfyAll(
					ContainSubstring("func AnyIoReadCloser() io.ReadCloser"),
				)),
				HaveKeyWithValue("map_of_string_to_interface", SatisfyAll(
					ContainSubstring("func AnyMapOfStringToInterface() map[string]interface{}"),
				)),
				HaveKeyWithValue("time_time", SatisfyAll(
					ContainSubstring("time \"time\""),
					ContainSubstring("func AnyTimeTime() time.Time"),
				)),
				HaveKeyWithValue("recv_chan_of_string", SatisfyAll(
					ContainSubstring("func AnyRecvChanOfString() <-chan string"),
				)),
				HaveKeyWithValue("send_chan_of_error", SatisfyAll(
					ContainSubstring("func AnySendChanOfError() chan<- error"),
				)),
				HaveKeyWithValue("map_of_int_to_int", SatisfyAll(
					ContainSubstring("func AnyMapOfIntToInt() map[int]int"),
				)),
				HaveKeyWithValue("map_of_http_file_to_http_file", SatisfyAll(
					ContainSubstring("http \"net/http\""),
					Not(MatchRegexp("http \"net/http\"\\s+http \"net/http\"")),
				)),
				HaveKeyWithValue("map_of_string_to_empty_unnamed_struct", SatisfyAll(
					ContainSubstring("func AnyMapOfStringToEmptyUnnamedStruct() map[string]struct{}"),
				)),
			))
		})
	})
})
