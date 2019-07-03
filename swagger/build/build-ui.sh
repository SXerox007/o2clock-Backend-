set -o errexit
set -o nounset
set -o pipefail

source "build/lib/init.sh"

if ! which go-bindata > /dev/null 2>&1 ; then
  echo "Cannot find go-bindata. Install with \"go get github.com/jteeuwen/go-bindata/...\""
  exit 1
fi

readonly TMP_DATAFILE="/tmp/datafile.go"
readonly SWAGGER_SRC="third_party/swagger-ui/..."
readonly SWAGGER_PKG="swagger"

function kube::hack::build_ui() {
  local pkg="$1"
  local src="$2"
  local output_file="pkg/ui/data/${pkg}/datafile.go"

  go-bindata -nocompress -o "${output_file}" -prefix ${PWD} -pkg "${pkg}" "${src}"

  local year=$(date +%Y)
  echo -e "// generated by hack/build-ui.sh; DO NOT EDIT\n" >> "${TMP_DATAFILE}"
  cat "${output_file}" >> "${TMP_DATAFILE}"

  gofmt -s -w "${TMP_DATAFILE}"

  mv "${TMP_DATAFILE}" "${output_file}"
}

kube::hack::build_ui "${SWAGGER_PKG}" "${SWAGGER_SRC}"
