// Auto-generated by avdl-compiler v1.3.16 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/user.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type Tracker struct {
	Tracker UID  `codec:"tracker" json:"tracker"`
	Status  int  `codec:"status" json:"status"`
	MTime   Time `codec:"mTime" json:"mTime"`
}

func (o Tracker) DeepCopy() Tracker {
	return Tracker{
		Tracker: o.Tracker.DeepCopy(),
		Status:  o.Status,
		MTime:   o.MTime.DeepCopy(),
	}
}

type TrackProof struct {
	ProofType string `codec:"proofType" json:"proofType"`
	ProofName string `codec:"proofName" json:"proofName"`
	IdString  string `codec:"idString" json:"idString"`
}

func (o TrackProof) DeepCopy() TrackProof {
	return TrackProof{
		ProofType: o.ProofType,
		ProofName: o.ProofName,
		IdString:  o.IdString,
	}
}

type WebProof struct {
	Hostname  string   `codec:"hostname" json:"hostname"`
	Protocols []string `codec:"protocols" json:"protocols"`
}

func (o WebProof) DeepCopy() WebProof {
	return WebProof{
		Hostname: o.Hostname,
		Protocols: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Protocols),
	}
}

type Proofs struct {
	Social     []TrackProof `codec:"social" json:"social"`
	Web        []WebProof   `codec:"web" json:"web"`
	PublicKeys []PublicKey  `codec:"publicKeys" json:"publicKeys"`
}

func (o Proofs) DeepCopy() Proofs {
	return Proofs{
		Social: (func(x []TrackProof) []TrackProof {
			var ret []TrackProof
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Social),
		Web: (func(x []WebProof) []WebProof {
			var ret []WebProof
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Web),
		PublicKeys: (func(x []PublicKey) []PublicKey {
			var ret []PublicKey
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.PublicKeys),
	}
}

type UserSummary struct {
	Uid          UID    `codec:"uid" json:"uid"`
	Username     string `codec:"username" json:"username"`
	Thumbnail    string `codec:"thumbnail" json:"thumbnail"`
	IdVersion    int    `codec:"idVersion" json:"idVersion"`
	FullName     string `codec:"fullName" json:"fullName"`
	Bio          string `codec:"bio" json:"bio"`
	Proofs       Proofs `codec:"proofs" json:"proofs"`
	SigIDDisplay string `codec:"sigIDDisplay" json:"sigIDDisplay"`
	TrackTime    Time   `codec:"trackTime" json:"trackTime"`
}

func (o UserSummary) DeepCopy() UserSummary {
	return UserSummary{
		Uid:          o.Uid.DeepCopy(),
		Username:     o.Username,
		Thumbnail:    o.Thumbnail,
		IdVersion:    o.IdVersion,
		FullName:     o.FullName,
		Bio:          o.Bio,
		Proofs:       o.Proofs.DeepCopy(),
		SigIDDisplay: o.SigIDDisplay,
		TrackTime:    o.TrackTime.DeepCopy(),
	}
}

type Email struct {
	Email      string `codec:"email" json:"email"`
	IsVerified bool   `codec:"isVerified" json:"isVerified"`
}

func (o Email) DeepCopy() Email {
	return Email{
		Email:      o.Email,
		IsVerified: o.IsVerified,
	}
}

type UserSettings struct {
	Emails []Email `codec:"emails" json:"emails"`
}

func (o UserSettings) DeepCopy() UserSettings {
	return UserSettings{
		Emails: (func(x []Email) []Email {
			var ret []Email
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Emails),
	}
}

type SearchComponent struct {
	Key   string  `codec:"key" json:"key"`
	Value string  `codec:"value" json:"value"`
	Score float64 `codec:"score" json:"score"`
}

func (o SearchComponent) DeepCopy() SearchComponent {
	return SearchComponent{
		Key:   o.Key,
		Value: o.Value,
		Score: o.Score,
	}
}

type SearchResult struct {
	Uid        UID               `codec:"uid" json:"uid"`
	Username   string            `codec:"username" json:"username"`
	Components []SearchComponent `codec:"components" json:"components"`
	Score      float64           `codec:"score" json:"score"`
}

func (o SearchResult) DeepCopy() SearchResult {
	return SearchResult{
		Uid:      o.Uid.DeepCopy(),
		Username: o.Username,
		Components: (func(x []SearchComponent) []SearchComponent {
			var ret []SearchComponent
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Components),
		Score: o.Score,
	}
}

type UserSummary2 struct {
	Uid        UID    `codec:"uid" json:"uid"`
	Username   string `codec:"username" json:"username"`
	Thumbnail  string `codec:"thumbnail" json:"thumbnail"`
	FullName   string `codec:"fullName" json:"fullName"`
	IsFollower bool   `codec:"isFollower" json:"isFollower"`
	IsFollowee bool   `codec:"isFollowee" json:"isFollowee"`
}

func (o UserSummary2) DeepCopy() UserSummary2 {
	return UserSummary2{
		Uid:        o.Uid.DeepCopy(),
		Username:   o.Username,
		Thumbnail:  o.Thumbnail,
		FullName:   o.FullName,
		IsFollower: o.IsFollower,
		IsFollowee: o.IsFollowee,
	}
}

type UserSummary2Set struct {
	Users   []UserSummary2 `codec:"users" json:"users"`
	Time    Time           `codec:"time" json:"time"`
	Version int            `codec:"version" json:"version"`
}

func (o UserSummary2Set) DeepCopy() UserSummary2Set {
	return UserSummary2Set{
		Users: (func(x []UserSummary2) []UserSummary2 {
			var ret []UserSummary2
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Users),
		Time:    o.Time.DeepCopy(),
		Version: o.Version,
	}
}

type ListTrackersArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
	Uid       UID `codec:"uid" json:"uid"`
}

func (o ListTrackersArg) DeepCopy() ListTrackersArg {
	return ListTrackersArg{
		SessionID: o.SessionID,
		Uid:       o.Uid.DeepCopy(),
	}
}

type ListTrackersByNameArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Username  string `codec:"username" json:"username"`
}

func (o ListTrackersByNameArg) DeepCopy() ListTrackersByNameArg {
	return ListTrackersByNameArg{
		SessionID: o.SessionID,
		Username:  o.Username,
	}
}

type ListTrackersSelfArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

func (o ListTrackersSelfArg) DeepCopy() ListTrackersSelfArg {
	return ListTrackersSelfArg{
		SessionID: o.SessionID,
	}
}

type LoadUncheckedUserSummariesArg struct {
	SessionID int   `codec:"sessionID" json:"sessionID"`
	Uids      []UID `codec:"uids" json:"uids"`
}

func (o LoadUncheckedUserSummariesArg) DeepCopy() LoadUncheckedUserSummariesArg {
	return LoadUncheckedUserSummariesArg{
		SessionID: o.SessionID,
		Uids: (func(x []UID) []UID {
			var ret []UID
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Uids),
	}
}

type LoadUserArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
	Uid       UID `codec:"uid" json:"uid"`
}

func (o LoadUserArg) DeepCopy() LoadUserArg {
	return LoadUserArg{
		SessionID: o.SessionID,
		Uid:       o.Uid.DeepCopy(),
	}
}

type LoadUserByNameArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Username  string `codec:"username" json:"username"`
}

func (o LoadUserByNameArg) DeepCopy() LoadUserByNameArg {
	return LoadUserByNameArg{
		SessionID: o.SessionID,
		Username:  o.Username,
	}
}

type LoadUserPlusKeysArg struct {
	SessionID  int `codec:"sessionID" json:"sessionID"`
	Uid        UID `codec:"uid" json:"uid"`
	PollForKID KID `codec:"pollForKID" json:"pollForKID"`
}

func (o LoadUserPlusKeysArg) DeepCopy() LoadUserPlusKeysArg {
	return LoadUserPlusKeysArg{
		SessionID:  o.SessionID,
		Uid:        o.Uid.DeepCopy(),
		PollForKID: o.PollForKID.DeepCopy(),
	}
}

type LoadPublicKeysArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
	Uid       UID `codec:"uid" json:"uid"`
}

func (o LoadPublicKeysArg) DeepCopy() LoadPublicKeysArg {
	return LoadPublicKeysArg{
		SessionID: o.SessionID,
		Uid:       o.Uid.DeepCopy(),
	}
}

type LoadMyPublicKeysArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

func (o LoadMyPublicKeysArg) DeepCopy() LoadMyPublicKeysArg {
	return LoadMyPublicKeysArg{
		SessionID: o.SessionID,
	}
}

type LoadMySettingsArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

func (o LoadMySettingsArg) DeepCopy() LoadMySettingsArg {
	return LoadMySettingsArg{
		SessionID: o.SessionID,
	}
}

type ListTrackingArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Filter    string `codec:"filter" json:"filter"`
	Assertion string `codec:"assertion" json:"assertion"`
}

func (o ListTrackingArg) DeepCopy() ListTrackingArg {
	return ListTrackingArg{
		SessionID: o.SessionID,
		Filter:    o.Filter,
		Assertion: o.Assertion,
	}
}

type ListTrackingJSONArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Filter    string `codec:"filter" json:"filter"`
	Verbose   bool   `codec:"verbose" json:"verbose"`
	Assertion string `codec:"assertion" json:"assertion"`
}

func (o ListTrackingJSONArg) DeepCopy() ListTrackingJSONArg {
	return ListTrackingJSONArg{
		SessionID: o.SessionID,
		Filter:    o.Filter,
		Verbose:   o.Verbose,
		Assertion: o.Assertion,
	}
}

type SearchArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Query     string `codec:"query" json:"query"`
}

func (o SearchArg) DeepCopy() SearchArg {
	return SearchArg{
		SessionID: o.SessionID,
		Query:     o.Query,
	}
}

type LoadAllPublicKeysUnverifiedArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
	Uid       UID `codec:"uid" json:"uid"`
}

func (o LoadAllPublicKeysUnverifiedArg) DeepCopy() LoadAllPublicKeysUnverifiedArg {
	return LoadAllPublicKeysUnverifiedArg{
		SessionID: o.SessionID,
		Uid:       o.Uid.DeepCopy(),
	}
}

type ListTrackers2Arg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Assertion string `codec:"assertion" json:"assertion"`
	Reverse   bool   `codec:"reverse" json:"reverse"`
}

func (o ListTrackers2Arg) DeepCopy() ListTrackers2Arg {
	return ListTrackers2Arg{
		SessionID: o.SessionID,
		Assertion: o.Assertion,
		Reverse:   o.Reverse,
	}
}

type ProfileEditArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	FullName  string `codec:"fullName" json:"fullName"`
	Location  string `codec:"location" json:"location"`
	Bio       string `codec:"bio" json:"bio"`
}

func (o ProfileEditArg) DeepCopy() ProfileEditArg {
	return ProfileEditArg{
		SessionID: o.SessionID,
		FullName:  o.FullName,
		Location:  o.Location,
		Bio:       o.Bio,
	}
}

type UserInterface interface {
	ListTrackers(context.Context, ListTrackersArg) ([]Tracker, error)
	ListTrackersByName(context.Context, ListTrackersByNameArg) ([]Tracker, error)
	ListTrackersSelf(context.Context, int) ([]Tracker, error)
	// Load user summaries for the supplied uids.
	// They are "unchecked" in that the client is not verifying the info from the server.
	// If len(uids) > 500, the first 500 will be returned.
	LoadUncheckedUserSummaries(context.Context, LoadUncheckedUserSummariesArg) ([]UserSummary, error)
	// Load a user from the server.
	LoadUser(context.Context, LoadUserArg) (User, error)
	LoadUserByName(context.Context, LoadUserByNameArg) (User, error)
	// Load a user + device keys from the server.
	LoadUserPlusKeys(context.Context, LoadUserPlusKeysArg) (UserPlusKeys, error)
	// Load public keys for a user.
	LoadPublicKeys(context.Context, LoadPublicKeysArg) ([]PublicKey, error)
	// Load my public keys (for logged in user).
	LoadMyPublicKeys(context.Context, int) ([]PublicKey, error)
	// Load user settings (for logged in user).
	LoadMySettings(context.Context, int) (UserSettings, error)
	// The list-tracking functions get verified data from the tracking statements
	// in the user's sigchain.
	//
	// If assertion is empty, it will use the current logged in user.
	ListTracking(context.Context, ListTrackingArg) ([]UserSummary, error)
	ListTrackingJSON(context.Context, ListTrackingJSONArg) (string, error)
	// Search for users who match a given query.
	Search(context.Context, SearchArg) ([]SearchResult, error)
	// Load all the user's public keys (even those in reset key families)
	// from the server with no verification
	LoadAllPublicKeysUnverified(context.Context, LoadAllPublicKeysUnverifiedArg) ([]PublicKey, error)
	ListTrackers2(context.Context, ListTrackers2Arg) (UserSummary2Set, error)
	ProfileEdit(context.Context, ProfileEditArg) error
}

func UserProtocol(i UserInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.user",
		Methods: map[string]rpc.ServeHandlerDescription{
			"listTrackers": {
				MakeArg: func() interface{} {
					ret := make([]ListTrackersArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ListTrackersArg)
					if !ok {
						err = rpc.NewTypeError((*[]ListTrackersArg)(nil), args)
						return
					}
					ret, err = i.ListTrackers(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"listTrackersByName": {
				MakeArg: func() interface{} {
					ret := make([]ListTrackersByNameArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ListTrackersByNameArg)
					if !ok {
						err = rpc.NewTypeError((*[]ListTrackersByNameArg)(nil), args)
						return
					}
					ret, err = i.ListTrackersByName(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"listTrackersSelf": {
				MakeArg: func() interface{} {
					ret := make([]ListTrackersSelfArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ListTrackersSelfArg)
					if !ok {
						err = rpc.NewTypeError((*[]ListTrackersSelfArg)(nil), args)
						return
					}
					ret, err = i.ListTrackersSelf(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"loadUncheckedUserSummaries": {
				MakeArg: func() interface{} {
					ret := make([]LoadUncheckedUserSummariesArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]LoadUncheckedUserSummariesArg)
					if !ok {
						err = rpc.NewTypeError((*[]LoadUncheckedUserSummariesArg)(nil), args)
						return
					}
					ret, err = i.LoadUncheckedUserSummaries(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"loadUser": {
				MakeArg: func() interface{} {
					ret := make([]LoadUserArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]LoadUserArg)
					if !ok {
						err = rpc.NewTypeError((*[]LoadUserArg)(nil), args)
						return
					}
					ret, err = i.LoadUser(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"loadUserByName": {
				MakeArg: func() interface{} {
					ret := make([]LoadUserByNameArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]LoadUserByNameArg)
					if !ok {
						err = rpc.NewTypeError((*[]LoadUserByNameArg)(nil), args)
						return
					}
					ret, err = i.LoadUserByName(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"loadUserPlusKeys": {
				MakeArg: func() interface{} {
					ret := make([]LoadUserPlusKeysArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]LoadUserPlusKeysArg)
					if !ok {
						err = rpc.NewTypeError((*[]LoadUserPlusKeysArg)(nil), args)
						return
					}
					ret, err = i.LoadUserPlusKeys(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"loadPublicKeys": {
				MakeArg: func() interface{} {
					ret := make([]LoadPublicKeysArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]LoadPublicKeysArg)
					if !ok {
						err = rpc.NewTypeError((*[]LoadPublicKeysArg)(nil), args)
						return
					}
					ret, err = i.LoadPublicKeys(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"loadMyPublicKeys": {
				MakeArg: func() interface{} {
					ret := make([]LoadMyPublicKeysArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]LoadMyPublicKeysArg)
					if !ok {
						err = rpc.NewTypeError((*[]LoadMyPublicKeysArg)(nil), args)
						return
					}
					ret, err = i.LoadMyPublicKeys(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"loadMySettings": {
				MakeArg: func() interface{} {
					ret := make([]LoadMySettingsArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]LoadMySettingsArg)
					if !ok {
						err = rpc.NewTypeError((*[]LoadMySettingsArg)(nil), args)
						return
					}
					ret, err = i.LoadMySettings(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"listTracking": {
				MakeArg: func() interface{} {
					ret := make([]ListTrackingArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ListTrackingArg)
					if !ok {
						err = rpc.NewTypeError((*[]ListTrackingArg)(nil), args)
						return
					}
					ret, err = i.ListTracking(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"listTrackingJSON": {
				MakeArg: func() interface{} {
					ret := make([]ListTrackingJSONArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ListTrackingJSONArg)
					if !ok {
						err = rpc.NewTypeError((*[]ListTrackingJSONArg)(nil), args)
						return
					}
					ret, err = i.ListTrackingJSON(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"search": {
				MakeArg: func() interface{} {
					ret := make([]SearchArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SearchArg)
					if !ok {
						err = rpc.NewTypeError((*[]SearchArg)(nil), args)
						return
					}
					ret, err = i.Search(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"loadAllPublicKeysUnverified": {
				MakeArg: func() interface{} {
					ret := make([]LoadAllPublicKeysUnverifiedArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]LoadAllPublicKeysUnverifiedArg)
					if !ok {
						err = rpc.NewTypeError((*[]LoadAllPublicKeysUnverifiedArg)(nil), args)
						return
					}
					ret, err = i.LoadAllPublicKeysUnverified(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"listTrackers2": {
				MakeArg: func() interface{} {
					ret := make([]ListTrackers2Arg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ListTrackers2Arg)
					if !ok {
						err = rpc.NewTypeError((*[]ListTrackers2Arg)(nil), args)
						return
					}
					ret, err = i.ListTrackers2(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"profileEdit": {
				MakeArg: func() interface{} {
					ret := make([]ProfileEditArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ProfileEditArg)
					if !ok {
						err = rpc.NewTypeError((*[]ProfileEditArg)(nil), args)
						return
					}
					err = i.ProfileEdit(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type UserClient struct {
	Cli rpc.GenericClient
}

func (c UserClient) ListTrackers(ctx context.Context, __arg ListTrackersArg) (res []Tracker, err error) {
	err = c.Cli.Call(ctx, "keybase.1.user.listTrackers", []interface{}{__arg}, &res)
	return
}

func (c UserClient) ListTrackersByName(ctx context.Context, __arg ListTrackersByNameArg) (res []Tracker, err error) {
	err = c.Cli.Call(ctx, "keybase.1.user.listTrackersByName", []interface{}{__arg}, &res)
	return
}

func (c UserClient) ListTrackersSelf(ctx context.Context, sessionID int) (res []Tracker, err error) {
	__arg := ListTrackersSelfArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.user.listTrackersSelf", []interface{}{__arg}, &res)
	return
}

// Load user summaries for the supplied uids.
// They are "unchecked" in that the client is not verifying the info from the server.
// If len(uids) > 500, the first 500 will be returned.
func (c UserClient) LoadUncheckedUserSummaries(ctx context.Context, __arg LoadUncheckedUserSummariesArg) (res []UserSummary, err error) {
	err = c.Cli.Call(ctx, "keybase.1.user.loadUncheckedUserSummaries", []interface{}{__arg}, &res)
	return
}

// Load a user from the server.
func (c UserClient) LoadUser(ctx context.Context, __arg LoadUserArg) (res User, err error) {
	err = c.Cli.Call(ctx, "keybase.1.user.loadUser", []interface{}{__arg}, &res)
	return
}

func (c UserClient) LoadUserByName(ctx context.Context, __arg LoadUserByNameArg) (res User, err error) {
	err = c.Cli.Call(ctx, "keybase.1.user.loadUserByName", []interface{}{__arg}, &res)
	return
}

// Load a user + device keys from the server.
func (c UserClient) LoadUserPlusKeys(ctx context.Context, __arg LoadUserPlusKeysArg) (res UserPlusKeys, err error) {
	err = c.Cli.Call(ctx, "keybase.1.user.loadUserPlusKeys", []interface{}{__arg}, &res)
	return
}

// Load public keys for a user.
func (c UserClient) LoadPublicKeys(ctx context.Context, __arg LoadPublicKeysArg) (res []PublicKey, err error) {
	err = c.Cli.Call(ctx, "keybase.1.user.loadPublicKeys", []interface{}{__arg}, &res)
	return
}

// Load my public keys (for logged in user).
func (c UserClient) LoadMyPublicKeys(ctx context.Context, sessionID int) (res []PublicKey, err error) {
	__arg := LoadMyPublicKeysArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.user.loadMyPublicKeys", []interface{}{__arg}, &res)
	return
}

// Load user settings (for logged in user).
func (c UserClient) LoadMySettings(ctx context.Context, sessionID int) (res UserSettings, err error) {
	__arg := LoadMySettingsArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.user.loadMySettings", []interface{}{__arg}, &res)
	return
}

// The list-tracking functions get verified data from the tracking statements
// in the user's sigchain.
//
// If assertion is empty, it will use the current logged in user.
func (c UserClient) ListTracking(ctx context.Context, __arg ListTrackingArg) (res []UserSummary, err error) {
	err = c.Cli.Call(ctx, "keybase.1.user.listTracking", []interface{}{__arg}, &res)
	return
}

func (c UserClient) ListTrackingJSON(ctx context.Context, __arg ListTrackingJSONArg) (res string, err error) {
	err = c.Cli.Call(ctx, "keybase.1.user.listTrackingJSON", []interface{}{__arg}, &res)
	return
}

// Search for users who match a given query.
func (c UserClient) Search(ctx context.Context, __arg SearchArg) (res []SearchResult, err error) {
	err = c.Cli.Call(ctx, "keybase.1.user.search", []interface{}{__arg}, &res)
	return
}

// Load all the user's public keys (even those in reset key families)
// from the server with no verification
func (c UserClient) LoadAllPublicKeysUnverified(ctx context.Context, __arg LoadAllPublicKeysUnverifiedArg) (res []PublicKey, err error) {
	err = c.Cli.Call(ctx, "keybase.1.user.loadAllPublicKeysUnverified", []interface{}{__arg}, &res)
	return
}

func (c UserClient) ListTrackers2(ctx context.Context, __arg ListTrackers2Arg) (res UserSummary2Set, err error) {
	err = c.Cli.Call(ctx, "keybase.1.user.listTrackers2", []interface{}{__arg}, &res)
	return
}

func (c UserClient) ProfileEdit(ctx context.Context, __arg ProfileEditArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.user.profileEdit", []interface{}{__arg}, nil)
	return
}
