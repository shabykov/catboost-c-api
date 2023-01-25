FROM --platform=$BUILDPLATFORM golang:1.16.3-buster AS build

ARG TARGETPLATFORM
ARG CATBOOST_SRC=/app/catboost_src
ARG CATBOOST_BIN=/app/catboost_bin

RUN echo "Platform=$TARGETPLATFORM"
RUN echo "Catboost src=$CATBOOST_SRC"
RUN echo "Catboost bin=$CATBOOST_BIN"

COPY catboost $CATBOOST_SRC

RUN mkdir -p $CATBOOST_BIN

RUN apt-get update && apt-get install git libc6-dev -y && rm -rf /var/lib/apt/lists/* \ 
	# && git clone -b v1.1.1 https://github.com/catboost/catboost.git $CATBOOST_SRC \
	&& cd $CATBOOST_SRC/catboost/libs/model_interface \
	&& ../../../ya make -r . \
	&& cp libcatboostmodel.* c_api.* model_calcer_wrapper.h wrapped_calcer.h $CATBOOST_BIN/ \
        && cd && rm -rf $CATBOOST_SRC \
	&& ls $CATBOOST_BIN

FROM alpine:latest

ARG CATBOOST_BIN=/app/catboost_bin

COPY --from=build $CATBOOST_BIN /bin
