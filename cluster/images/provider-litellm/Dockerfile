FROM alpine:3.23.4
RUN apk --no-cache add ca-certificates bash

ARG TARGETOS
ARG TARGETARCH

ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

ENV USER_ID=65532

# Setup Terraform environment

## Provider-dependent configuration
ARG TERRAFORM_VERSION
ARG TERRAFORM_PROVIDER_SOURCE
ARG TERRAFORM_PROVIDER_VERSION
ARG TERRAFORM_PROVIDER_DOWNLOAD_NAME
ARG TERRAFORM_NATIVE_PROVIDER_BINARY

## End of - Provider-dependent configuration

# OpenTofu resolves an unqualified provider source against its own registry
# host, so the in-image mirror is keyed by registry.opentofu.org.
ENV PLUGIN_DIR=/terraform/provider-mirror/registry.opentofu.org/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_PROVIDER_VERSION}/${TARGETOS}_${TARGETARCH}
ENV TF_CLI_CONFIG_FILE=/terraform/.terraformrc
ENV TF_FORK=0

RUN mkdir -p ${PLUGIN_DIR}

# OpenTofu, installed as `terraform` because that is the executable upjet
# invokes. Terraform itself is BSL licensed from 1.6 and cannot be redistributed
# in this image; OpenTofu is the MPL-2.0 fork.
ADD https://github.com/opentofu/opentofu/releases/download/v${TERRAFORM_VERSION}/tofu_${TERRAFORM_VERSION}_${TARGETOS}_${TARGETARCH}.zip /tmp/tofu.zip
ADD terraformrc.hcl ${TF_CLI_CONFIG_FILE}

# The provider is built from an upstream branch that publishes no release
# artifact, so the binary is staged into the build context by the image Makefile
# rather than downloaded. See hack/build-terraform-provider.sh.
COPY tf-provider/${TARGETOS}_${TARGETARCH}/${TERRAFORM_NATIVE_PROVIDER_BINARY} ${PLUGIN_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}

RUN unzip /tmp/tofu.zip -d /tmp/tofu \
  && mv /tmp/tofu/tofu /usr/local/bin/terraform \
  && chmod +x /usr/local/bin/terraform \
  && rm -rf /tmp/tofu /tmp/tofu.zip \
  && chmod +x ${PLUGIN_DIR}/* \
  && chown -R ${USER_ID}:${USER_ID} /terraform
# End of - Setup Terraform environment

# Provider controller needs these environment variable at runtime
ENV TERRAFORM_VERSION=${TERRAFORM_VERSION}
ENV TERRAFORM_PROVIDER_SOURCE=${TERRAFORM_PROVIDER_SOURCE}
ENV TERRAFORM_PROVIDER_VERSION=${TERRAFORM_PROVIDER_VERSION}
ENV TERRAFORM_NATIVE_PROVIDER_PATH=${PLUGIN_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
