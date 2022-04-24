package http

const VersionTypeName = "Version"
const VersionTypeFullName = "mojo.http.Version"

func (x *Version) IsEmpty() bool {
    return x == nil || x.Major == 0
}
