package portal

type Mountable interface {
	MountPortal(*Portal)
}

func MountPortalToMutator(portal *Portal, x any) {
	if mountable, isMountable := x.(Mountable); isMountable {
		mountable.MountPortal(portal)
	}
}
