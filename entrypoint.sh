#!/bin/sh

set -e

if [ ! -z "${PLUGIN_DEBUG}" ]; then
	set -x
fi

if [ ! -z "${PLUGIN_PAUSE}" ]; then 
 	sleep 100000
fi

if [ ! -z "${PLUGIN_PROXY}" ]; then
	export http_proxy=${PLUGIN_PROXY}
	export https_proxy=${PLUGIN_PROXY}
	export all_proxy=${PLUGIN_PROXY}
	export no_proxy=localhost,127.0.0.1/8
	echo "http proxy done"
fi

export PLUGIN_VERSIONFILE=${PLUGIN_VERSIONFILE:-VERSION}

if [ -z "${PLUGIN_RELEASE}" ]; then
	if [ -f ${PLUGIN_VERSIONFILE} ]; then
		export PLUGIN_RELEASE=$(cat ${PLUGIN_VERSIONFILE})
	else
		echo "${PLUGIN_VERSIONFILE} does not exist"
		exit 1
	fi
fi

exec "$@"