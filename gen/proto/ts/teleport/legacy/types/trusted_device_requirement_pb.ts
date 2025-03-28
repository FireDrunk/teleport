/* eslint-disable */
// @generated by protobuf-ts 2.9.3 with parameter eslint_disable,add_pb_suffix,server_grpc1,ts_nocheck
// @generated from protobuf file "teleport/legacy/types/trusted_device_requirement.proto" (package "types", syntax proto3)
// tslint:disable
// @ts-nocheck
//
// Copyright 2024 Gravitational, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
/**
 * TrustedDeviceRequirement indicates whether access may be hindered by the lack
 * of a trusted device.
 *
 * @generated from protobuf enum types.TrustedDeviceRequirement
 */
export enum TrustedDeviceRequirement {
    /**
     * Device requirement not determined.
     * Does not mean that a device is not required, only that the necessary data
     * was not considered.
     *
     * @generated from protobuf enum value: TRUSTED_DEVICE_REQUIREMENT_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * Trusted device not required.
     *
     * @generated from protobuf enum value: TRUSTED_DEVICE_REQUIREMENT_NOT_REQUIRED = 1;
     */
    NOT_REQUIRED = 1,
    /**
     * Trusted device required by either cluster mode or user roles.
     *
     * @generated from protobuf enum value: TRUSTED_DEVICE_REQUIREMENT_REQUIRED = 2;
     */
    REQUIRED = 2
}
