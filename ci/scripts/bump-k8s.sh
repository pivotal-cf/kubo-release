#!/bin/bash
set -exu -o pipefail

source git-pks-kubo-release-ci/ci/scripts/lib/generate-pr.sh

pr_release() {
  version="$1"
  git_repo="$2"
  script_name="$3"
  base_branch="$4"

  concourse_base_name="git-${git_repo}"

  cp -r "${concourse_base_name}/." "${concourse_base_name}-output"
  pushd "${concourse_base_name}-output"

  ../git-pks-kubo-release-ci/ci/scripts/$script_name $version $(pwd)

  if [ -n "$(git status --porcelain)" ]; then
    cat <<EOF > "config/private.yml"
blobstore:
  options:
    access_key_id: ${ACCESS_KEY_ID}
    secret_access_key: ${SECRET_ACCESS_KEY}
EOF
    bosh upload-blobs
    generate_pull_request "kubernetes" "$version" "${git_repo}" "${base_branch}"
  else
    echo "Kubernetes version is already up-to-date"
  fi

  popd
}

main() {

  if [ "${REPO:-}" == "windows" ]; then
    git_repo="pks-kubo-release-windows"
    script_name="download_k8s_binaries_windows"
  else
    git_repo="pks-kubo-release"
    if [ "false" == "$USE_COMMON_CORE" ]; then
      script_name="download_k8s_binaries_google"
    else
      script_name="download_k8s_binaries_common_core"
    fi
  fi

  # BINARY_DIRECTORY should be declared in the pipeline via params
  version=$(cat "$PWD/$BINARY_DIRECTORY/version")
  pr_release "$version" "$git_repo" "$script_name" "${BASE_BRANCH}"
}

main $@