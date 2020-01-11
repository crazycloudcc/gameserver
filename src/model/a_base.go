package model

/*
 * module model.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

// Redis Keys
const (
	// STRINGversion   string = "version"   // version
	// STRINGupdateurl string = "updateURL" // updateURL

	// remote global
	INCRseeduid  string = "incr_game001_uid_seed" // global game001 uid seed
	HASHuidtoken string = "hash_uid_token"        // uid => TokenInfo

	// local
	HASHuiduser string = "hash_uid_user" // uid => UserInfo
)

/************************************************************************/
// export functions.
/************************************************************************/

/************************************************************************/
// moudule functions.
/************************************************************************/

// InitFromRedis TODO
func InitFromRedis() {
	SeedUIDInit()
}

/************************************************************************/
// unit tests.
/************************************************************************/
