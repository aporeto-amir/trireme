// Package policy describes a generic interface for retrieving policies.
// Different implementations are possible for environments such as Kubernetes,
// Mesos or other custom environments. An implementation has to provide
// a method for retrieving policy based on the metadata associated with the container
// and deleting the policy when the container dies. It is up to the implementation
// to decide how to generate the policy.
// The package also defines the basic data structure for communicating policy
// information. The implementations are responsible for providing all the necessary
// data.
package policy

// A RuntimeReader allows to get the specific parameters stored in the Runtime
type RuntimeReader interface {

	// Pid returns the Pid of the Runtime.
	Pid() int

	// Name returns the process name of the Runtime.
	Name() string

	// Tag retuns the value of the given tag.
	Tag(string) (string, bool)

	// Tags returns a copy of the list of the tags.
	Tags() *TagsMap

	// DefaultIPAddress retutns the default IP address.
	DefaultIPAddress() (string, bool)

	// IPAddresses returns a copy of all the IP addresses.
	IPAddresses() *IPMap
}

// InfoInteractor is the interface for setting up policy before providing to trireme
type InfoInteractor interface {

	// Clone returns a copy of the policy
	Clone() *PUPolicy

	// IngressACLs returns a copy of IPRuleList
	IngressACLs() *IPRuleList

	// SetIngressACLs adds ingress rules
	SetIngressACLs(r *IPRuleList)

	// EgressACLs returns a copy of IPRuleList
	EgressACLs() *IPRuleList

	// SetEgressACLs adds ingress rules
	SetEgressACLs(r *IPRuleList)

	// ReceiverRules returns a copy of TagSelectorList
	ReceiverRules() *TagSelectorList

	// AddReceiverRules adds a receiver rule
	AddReceiverRules(t *TagSelector)

	// TransmitterRules returns a copy of TagSelectorList
	TransmitterRules() *TagSelectorList

	// AddTransmitterRules adds a transmitter rule
	AddTransmitterRules(t *TagSelector)

	// PolicyTags returns a copy of PolicyTag(s)
	Indentity() *TagsMap

	// SetPolicyTags sets up policy tags
	SetIdentity(t *TagsMap)

	// AddPolicyTag adds a policy tag
	AddPolicyTag(k, v string)

	// DefaultIPAddress returns the default IP address for the processing unit
	DefaultIPAddress() (string, bool)
}
