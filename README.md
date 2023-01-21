# Device Service

Manages a list of known devices and their version assignments.

## API:

* `GET /` - Get the list of devices
* `GET /<mac>` - Get a device record by its mac address
* `POST /<mac>` - Update a device, sets the last updated time, current firmware and version
* `POST /<mac>/name?val=<name>` - Update the device name
* `POST /<mac>/firmware?val=<name>` - Update the device firmware
* `POST /<mac>/version?val=<name>` - Update the device firmware version constraint
