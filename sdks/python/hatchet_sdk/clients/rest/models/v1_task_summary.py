# coding: utf-8

"""
    Hatchet API

    The Hatchet API

    The version of the OpenAPI document: 1.0.0
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations

import json
import pprint
import re  # noqa: F401
from datetime import datetime
from typing import Any, ClassVar, Dict, List, Optional, Set

from pydantic import BaseModel, ConfigDict, Field, StrictInt, StrictStr
from typing_extensions import Annotated, Self

from hatchet_sdk.clients.rest.models.api_resource_meta import APIResourceMeta
from hatchet_sdk.clients.rest.models.v1_task_status import V1TaskStatus
from hatchet_sdk.clients.rest.models.v1_workflow_type import V1WorkflowType


class V1TaskSummary(BaseModel):
    """
    V1TaskSummary
    """  # noqa: E501

    metadata: APIResourceMeta
    action_id: Optional[StrictStr] = Field(
        default=None, description="The action ID of the task.", alias="actionId"
    )
    retry_count: Optional[StrictInt] = Field(
        default=None,
        description="The number of retries of the task.",
        alias="retryCount",
    )
    attempt: Optional[StrictInt] = Field(
        default=None, description="The attempt number of the task."
    )
    additional_metadata: Optional[Dict[str, Any]] = Field(
        default=None,
        description="Additional metadata for the task run.",
        alias="additionalMetadata",
    )
    children: Optional[List[V1TaskSummary]] = Field(
        default=None, description="The list of children tasks"
    )
    created_at: datetime = Field(
        description="The timestamp the task was created.", alias="createdAt"
    )
    display_name: StrictStr = Field(
        description="The display name of the task run.", alias="displayName"
    )
    duration: Optional[StrictInt] = Field(
        default=None, description="The duration of the task run, in milliseconds."
    )
    error_message: Optional[StrictStr] = Field(
        default=None,
        description="The error message of the task run (for the latest run)",
        alias="errorMessage",
    )
    finished_at: Optional[datetime] = Field(
        default=None,
        description="The timestamp the task run finished.",
        alias="finishedAt",
    )
    input: Dict[str, Any] = Field(description="The input of the task run.")
    num_spawned_children: StrictInt = Field(
        description="The number of spawned children tasks", alias="numSpawnedChildren"
    )
    output: Dict[str, Any] = Field(
        description="The output of the task run (for the latest run)"
    )
    status: V1TaskStatus
    started_at: Optional[datetime] = Field(
        default=None,
        description="The timestamp the task run started.",
        alias="startedAt",
    )
    step_id: Optional[
        Annotated[str, Field(min_length=36, strict=True, max_length=36)]
    ] = Field(default=None, description="The step ID of the task.", alias="stepId")
    task_external_id: Annotated[
        str, Field(min_length=36, strict=True, max_length=36)
    ] = Field(description="The external ID of the task.", alias="taskExternalId")
    task_id: StrictInt = Field(description="The ID of the task.", alias="taskId")
    task_inserted_at: datetime = Field(
        description="The timestamp the task was inserted.", alias="taskInsertedAt"
    )
    tenant_id: Annotated[str, Field(min_length=36, strict=True, max_length=36)] = Field(
        description="The ID of the tenant.", alias="tenantId"
    )
    type: V1WorkflowType = Field(
        description="The type of the workflow (whether it's a DAG or a task)"
    )
    workflow_id: StrictStr = Field(alias="workflowId")
    workflow_name: Optional[StrictStr] = Field(default=None, alias="workflowName")
    workflow_run_external_id: StrictStr = Field(
        description="The external ID of the workflow run", alias="workflowRunExternalId"
    )
    workflow_version_id: Optional[StrictStr] = Field(
        default=None,
        description="The version ID of the workflow",
        alias="workflowVersionId",
    )
    workflow_config: Optional[Dict[str, Any]] = Field(
        default=None, alias="workflowConfig"
    )
    parent_task_external_id: Optional[StrictStr] = Field(
        default=None,
        description="The external ID of the parent task.",
        alias="parentTaskExternalId",
    )
    __properties: ClassVar[List[str]] = [
        "metadata",
        "actionId",
        "retryCount",
        "attempt",
        "additionalMetadata",
        "children",
        "createdAt",
        "displayName",
        "duration",
        "errorMessage",
        "finishedAt",
        "input",
        "numSpawnedChildren",
        "output",
        "status",
        "startedAt",
        "stepId",
        "taskExternalId",
        "taskId",
        "taskInsertedAt",
        "tenantId",
        "type",
        "workflowId",
        "workflowName",
        "workflowRunExternalId",
        "workflowVersionId",
        "workflowConfig",
        "parentTaskExternalId",
    ]

    model_config = ConfigDict(
        populate_by_name=True,
        validate_assignment=True,
        protected_namespaces=(),
    )

    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Optional[Self]:
        """Create an instance of V1TaskSummary from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        excluded_fields: Set[str] = set([])

        _dict = self.model_dump(
            by_alias=True,
            exclude=excluded_fields,
            exclude_none=True,
        )
        # override the default output from pydantic by calling `to_dict()` of metadata
        if self.metadata:
            _dict["metadata"] = self.metadata.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in children (list)
        _items = []
        if self.children:
            for _item_children in self.children:
                if _item_children:
                    _items.append(_item_children.to_dict())
            _dict["children"] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of V1TaskSummary from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate(
            {
                "metadata": (
                    APIResourceMeta.from_dict(obj["metadata"])
                    if obj.get("metadata") is not None
                    else None
                ),
                "actionId": obj.get("actionId"),
                "retryCount": obj.get("retryCount"),
                "attempt": obj.get("attempt"),
                "additionalMetadata": obj.get("additionalMetadata"),
                "children": (
                    [V1TaskSummary.from_dict(_item) for _item in obj["children"]]
                    if obj.get("children") is not None
                    else None
                ),
                "createdAt": obj.get("createdAt"),
                "displayName": obj.get("displayName"),
                "duration": obj.get("duration"),
                "errorMessage": obj.get("errorMessage"),
                "finishedAt": obj.get("finishedAt"),
                "input": obj.get("input"),
                "numSpawnedChildren": obj.get("numSpawnedChildren"),
                "output": obj.get("output"),
                "status": obj.get("status"),
                "startedAt": obj.get("startedAt"),
                "stepId": obj.get("stepId"),
                "taskExternalId": obj.get("taskExternalId"),
                "taskId": obj.get("taskId"),
                "taskInsertedAt": obj.get("taskInsertedAt"),
                "tenantId": obj.get("tenantId"),
                "type": obj.get("type"),
                "workflowId": obj.get("workflowId"),
                "workflowName": obj.get("workflowName"),
                "workflowRunExternalId": obj.get("workflowRunExternalId"),
                "workflowVersionId": obj.get("workflowVersionId"),
                "workflowConfig": obj.get("workflowConfig"),
                "parentTaskExternalId": obj.get("parentTaskExternalId"),
            }
        )
        return _obj


# TODO: Rewrite to not use raise_errors
V1TaskSummary.model_rebuild(raise_errors=False)
