# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *
from .rec import *
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "example",
  "mod": "",
  "fqn": "pulumi_example",
  "classes": {
   "example::Rec": "Rec"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "example",
  "token": "pulumi:providers:example",
  "fqn": "pulumi_example",
  "class": "Provider"
 }
]
"""
)