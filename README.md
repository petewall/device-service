# Device Service

Serves devices

## TODO

* Port is a environment variable
* Connect with redis
* Implement API
  * `GET /` --> Gets the list of devices
  * `GET /<mac>` --> Gets details about a device
  * `GET /<mac>/history` --> Gets the full update history for a device
  * `POST /<mac>` w/ payload --> Update a device
* Implement unit tests
* Implement integration tests
* Create Dockerfile for distribution
* Create deployment files for k8s

Device payload:
{
  mac: xxxx
  currentFirmware: bootloader
  currentVersion: 0.0.1
  assignedFirmware: bootloader <-- indicates assigned firmware.
  assignedVersion: 0.0.1 <-- indicates pinned version. Null if no pinned version
  acceptsPrerelease: false <-- if true, will accept pre-release version updates
}
