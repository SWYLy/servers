/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package dao

import (
	"Swyl/servers/swyl-club-ms/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// @notice Root struct for other methods in club-dao-impl
type ClubDaoImpl struct {
	ctx 				context.Context
	mongCollection		*mongo.Collection
}

// @dev Constructor
func ClubDaoConstructor (ctx context.Context, mongCollection *mongo.Collection) ClubDao {
	return &ClubDaoImpl{
		ctx: ctx,
		mongCollection: mongCollection,
	}
}


// @notice Method of UserDaoImpl struct
// 
// @dev Creates a club
// 
// @param clubOwner *string
// 
// @return error
func (ci *ClubDaoImpl) CreateClub(clubOwner *string) error {
	return nil
}


// @notice Method of UserDaoImpl struct
// 
// @dev Gets a club at clubId
// 
// @param clubId *uint64
// 
// @return *models.Club
// 
// @return error
func (ci *ClubDaoImpl) GetClubAt(clubId *uint64) (*models.Club, error) {
	return nil, nil
}


// @notice Method of UserDaoImpl struct
// 
// @dev Gets a club at clubId
// 
// @param clubOwner *string
// 
// @return *models.Club
// 
// @return error
func (ci *ClubDaoImpl) GetClubOwnedBy(clubOwner *string) (*models.Club, error) {
	return nil, nil
}


// @notice Method of UserDaoImpl struct
// 
// @dev Updates a Club's Total_member
// 
// @param clubOwner *string
// 
// @param totalMember *uint64
func (ci *ClubDaoImpl) UpdateClubTotalMember(clubOwner *string, totalMember *uint64) error {
	return nil
}