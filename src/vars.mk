include $(dir $(lastword $(MAKEFILE_LIST))).env

REQUIRED_VARS := registry repo

$(foreach v,$(REQUIRED_VARS),\
  $(if $(value $(v)),,$(error Variable '$(v)' is required but empty)))

tag = v0.0.91-alpha
