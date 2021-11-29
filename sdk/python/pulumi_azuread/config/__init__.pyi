# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

clientCertificate: Optional[str]
"""
Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
"""

clientCertificatePassword: Optional[str]
"""
The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
Certificate
"""

clientCertificatePath: Optional[str]
"""
The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
Principal using a Client Certificate
"""

clientId: Optional[str]
"""
The Client ID which should be used for service principal authentication
"""

clientSecret: Optional[str]
"""
The application password to use when authenticating as a Service Principal using a Client Secret
"""

disableTerraformPartnerId: Optional[bool]
"""
Disable the Terraform Partner ID, which is used if a custom `partner_id` isn't specified
"""

environment: str
"""
The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
`usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`
"""

msiEndpoint: Optional[str]
"""
The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
"""

partnerId: Optional[str]
"""
A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
"""

tenantId: Optional[str]
"""
The Tenant ID which should be used. Works with all authentication methods except Managed Identity
"""

useCli: Optional[bool]
"""
Allow Azure CLI to be used for Authentication
"""

useMsi: bool
"""
Allow Managed Identity to be used for Authentication
"""

