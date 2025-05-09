// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/harness/gitness/registry/types"

	mock "github.com/stretchr/testify/mock"
)

// TagRepository is an autogenerated mock type for the TagRepository type
type TagRepository struct {
	mock.Mock
}

// CountAllArtifactsByParentID provides a mock function with given fields: ctx, parentID, registryIDs, search, latestVersion, packageTypes
func (_m *TagRepository) CountAllArtifactsByParentID(ctx context.Context, parentID int64, registryIDs *[]string, search string, latestVersion bool, packageTypes []string) (int64, error) {
	ret := _m.Called(ctx, parentID, registryIDs, search, latestVersion, packageTypes)

	if len(ret) == 0 {
		panic("no return value specified for CountAllArtifactsByParentID")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *[]string, string, bool, []string) (int64, error)); ok {
		return rf(ctx, parentID, registryIDs, search, latestVersion, packageTypes)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, *[]string, string, bool, []string) int64); ok {
		r0 = rf(ctx, parentID, registryIDs, search, latestVersion, packageTypes)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, *[]string, string, bool, []string) error); ok {
		r1 = rf(ctx, parentID, registryIDs, search, latestVersion, packageTypes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountAllArtifactsByRepo provides a mock function with given fields: ctx, parentID, repoKey, search, labels
func (_m *TagRepository) CountAllArtifactsByRepo(ctx context.Context, parentID int64, repoKey string, search string, labels []string) (int64, error) {
	ret := _m.Called(ctx, parentID, repoKey, search, labels)

	if len(ret) == 0 {
		panic("no return value specified for CountAllArtifactsByRepo")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string, []string) (int64, error)); ok {
		return rf(ctx, parentID, repoKey, search, labels)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string, []string) int64); ok {
		r0 = rf(ctx, parentID, repoKey, search, labels)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string, string, []string) error); ok {
		r1 = rf(ctx, parentID, repoKey, search, labels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountAllTagsByRepoAndImage provides a mock function with given fields: ctx, parentID, repoKey, image, search
func (_m *TagRepository) CountAllTagsByRepoAndImage(ctx context.Context, parentID int64, repoKey string, image string, search string) (int64, error) {
	ret := _m.Called(ctx, parentID, repoKey, image, search)

	if len(ret) == 0 {
		panic("no return value specified for CountAllTagsByRepoAndImage")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string, string) (int64, error)); ok {
		return rf(ctx, parentID, repoKey, image, search)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string, string) int64); ok {
		r0 = rf(ctx, parentID, repoKey, image, search)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string, string, string) error); ok {
		r1 = rf(ctx, parentID, repoKey, image, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrUpdate provides a mock function with given fields: ctx, t
func (_m *TagRepository) CreateOrUpdate(ctx context.Context, t *types.Tag) error {
	ret := _m.Called(ctx, t)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrUpdate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Tag) error); ok {
		r0 = rf(ctx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTag provides a mock function with given fields: ctx, registryID, imageName, name
func (_m *TagRepository) DeleteTag(ctx context.Context, registryID int64, imageName string, name string) error {
	ret := _m.Called(ctx, registryID, imageName, name)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string) error); ok {
		r0 = rf(ctx, registryID, imageName, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTagByManifestID provides a mock function with given fields: ctx, repoID, manifestID
func (_m *TagRepository) DeleteTagByManifestID(ctx context.Context, repoID int64, manifestID int64) (bool, error) {
	ret := _m.Called(ctx, repoID, manifestID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTagByManifestID")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) (bool, error)); ok {
		return rf(ctx, repoID, manifestID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) bool); ok {
		r0 = rf(ctx, repoID, manifestID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, repoID, manifestID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTagByName provides a mock function with given fields: ctx, repoID, name
func (_m *TagRepository) DeleteTagByName(ctx context.Context, repoID int64, name string) (bool, error) {
	ret := _m.Called(ctx, repoID, name)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTagByName")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) (bool, error)); ok {
		return rf(ctx, repoID, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) bool); ok {
		r0 = rf(ctx, repoID, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, repoID, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTagsByImageName provides a mock function with given fields: ctx, registryID, imageName
func (_m *TagRepository) DeleteTagsByImageName(ctx context.Context, registryID int64, imageName string) error {
	ret := _m.Called(ctx, registryID, imageName)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTagsByImageName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) error); ok {
		r0 = rf(ctx, registryID, imageName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindTag provides a mock function with given fields: ctx, repoID, imageName, name
func (_m *TagRepository) FindTag(ctx context.Context, repoID int64, imageName string, name string) (*types.Tag, error) {
	ret := _m.Called(ctx, repoID, imageName, name)

	if len(ret) == 0 {
		panic("no return value specified for FindTag")
	}

	var r0 *types.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string) (*types.Tag, error)); ok {
		return rf(ctx, repoID, imageName, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string) *types.Tag); ok {
		r0 = rf(ctx, repoID, imageName, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Tag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string, string) error); ok {
		r1 = rf(ctx, repoID, imageName, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllArtifactsByParentID provides a mock function with given fields: ctx, parentID, registryIDs, sortByField, sortByOrder, limit, offset, search, latestVersion, packageTypes
func (_m *TagRepository) GetAllArtifactsByParentID(ctx context.Context, parentID int64, registryIDs *[]string, sortByField string, sortByOrder string, limit int, offset int, search string, latestVersion bool, packageTypes []string) (*[]types.ArtifactMetadata, error) {
	ret := _m.Called(ctx, parentID, registryIDs, sortByField, sortByOrder, limit, offset, search, latestVersion, packageTypes)

	if len(ret) == 0 {
		panic("no return value specified for GetAllArtifactsByParentID")
	}

	var r0 *[]types.ArtifactMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *[]string, string, string, int, int, string, bool, []string) (*[]types.ArtifactMetadata, error)); ok {
		return rf(ctx, parentID, registryIDs, sortByField, sortByOrder, limit, offset, search, latestVersion, packageTypes)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, *[]string, string, string, int, int, string, bool, []string) *[]types.ArtifactMetadata); ok {
		r0 = rf(ctx, parentID, registryIDs, sortByField, sortByOrder, limit, offset, search, latestVersion, packageTypes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]types.ArtifactMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, *[]string, string, string, int, int, string, bool, []string) error); ok {
		r1 = rf(ctx, parentID, registryIDs, sortByField, sortByOrder, limit, offset, search, latestVersion, packageTypes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllArtifactsByRepo provides a mock function with given fields: ctx, parentID, repoKey, sortByField, sortByOrder, limit, offset, search, labels
func (_m *TagRepository) GetAllArtifactsByRepo(ctx context.Context, parentID int64, repoKey string, sortByField string, sortByOrder string, limit int, offset int, search string, labels []string) (*[]types.ArtifactMetadata, error) {
	ret := _m.Called(ctx, parentID, repoKey, sortByField, sortByOrder, limit, offset, search, labels)

	if len(ret) == 0 {
		panic("no return value specified for GetAllArtifactsByRepo")
	}

	var r0 *[]types.ArtifactMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string, string, int, int, string, []string) (*[]types.ArtifactMetadata, error)); ok {
		return rf(ctx, parentID, repoKey, sortByField, sortByOrder, limit, offset, search, labels)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string, string, int, int, string, []string) *[]types.ArtifactMetadata); ok {
		r0 = rf(ctx, parentID, repoKey, sortByField, sortByOrder, limit, offset, search, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]types.ArtifactMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string, string, string, int, int, string, []string) error); ok {
		r1 = rf(ctx, parentID, repoKey, sortByField, sortByOrder, limit, offset, search, labels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTagsByRepoAndImage provides a mock function with given fields: ctx, parentID, repoKey, image, sortByField, sortByOrder, limit, offset, search
func (_m *TagRepository) GetAllTagsByRepoAndImage(ctx context.Context, parentID int64, repoKey string, image string, sortByField string, sortByOrder string, limit int, offset int, search string) (*[]types.TagMetadata, error) {
	ret := _m.Called(ctx, parentID, repoKey, image, sortByField, sortByOrder, limit, offset, search)

	if len(ret) == 0 {
		panic("no return value specified for GetAllTagsByRepoAndImage")
	}

	var r0 *[]types.TagMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string, string, string, int, int, string) (*[]types.TagMetadata, error)); ok {
		return rf(ctx, parentID, repoKey, image, sortByField, sortByOrder, limit, offset, search)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string, string, string, int, int, string) *[]types.TagMetadata); ok {
		r0 = rf(ctx, parentID, repoKey, image, sortByField, sortByOrder, limit, offset, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]types.TagMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string, string, string, string, int, int, string) error); ok {
		r1 = rf(ctx, parentID, repoKey, image, sortByField, sortByOrder, limit, offset, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestTag provides a mock function with given fields: ctx, repoID, imageName
func (_m *TagRepository) GetLatestTag(ctx context.Context, repoID int64, imageName string) (*types.Tag, error) {
	ret := _m.Called(ctx, repoID, imageName)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestTag")
	}

	var r0 *types.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) (*types.Tag, error)); ok {
		return rf(ctx, repoID, imageName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) *types.Tag); ok {
		r0 = rf(ctx, repoID, imageName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Tag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, repoID, imageName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestTagMetadata provides a mock function with given fields: ctx, parentID, repoKey, imageName
func (_m *TagRepository) GetLatestTagMetadata(ctx context.Context, parentID int64, repoKey string, imageName string) (*types.ArtifactMetadata, error) {
	ret := _m.Called(ctx, parentID, repoKey, imageName)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestTagMetadata")
	}

	var r0 *types.ArtifactMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string) (*types.ArtifactMetadata, error)); ok {
		return rf(ctx, parentID, repoKey, imageName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string) *types.ArtifactMetadata); ok {
		r0 = rf(ctx, parentID, repoKey, imageName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ArtifactMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string, string) error); ok {
		r1 = rf(ctx, parentID, repoKey, imageName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestTagName provides a mock function with given fields: ctx, parentID, repoKey, imageName
func (_m *TagRepository) GetLatestTagName(ctx context.Context, parentID int64, repoKey string, imageName string) (string, error) {
	ret := _m.Called(ctx, parentID, repoKey, imageName)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestTagName")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string) (string, error)); ok {
		return rf(ctx, parentID, repoKey, imageName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string) string); ok {
		r0 = rf(ctx, parentID, repoKey, imageName)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string, string) error); ok {
		r1 = rf(ctx, parentID, repoKey, imageName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTagDetail provides a mock function with given fields: ctx, repoID, imageName, name
func (_m *TagRepository) GetTagDetail(ctx context.Context, repoID int64, imageName string, name string) (*types.TagDetail, error) {
	ret := _m.Called(ctx, repoID, imageName, name)

	if len(ret) == 0 {
		panic("no return value specified for GetTagDetail")
	}

	var r0 *types.TagDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string) (*types.TagDetail, error)); ok {
		return rf(ctx, repoID, imageName, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string) *types.TagDetail); ok {
		r0 = rf(ctx, repoID, imageName, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.TagDetail)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string, string) error); ok {
		r1 = rf(ctx, repoID, imageName, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTagMetadata provides a mock function with given fields: ctx, parentID, repoKey, imageName, name
func (_m *TagRepository) GetTagMetadata(ctx context.Context, parentID int64, repoKey string, imageName string, name string) (*types.TagMetadata, error) {
	ret := _m.Called(ctx, parentID, repoKey, imageName, name)

	if len(ret) == 0 {
		panic("no return value specified for GetTagMetadata")
	}

	var r0 *types.TagMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string, string) (*types.TagMetadata, error)); ok {
		return rf(ctx, parentID, repoKey, imageName, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string, string) *types.TagMetadata); ok {
		r0 = rf(ctx, parentID, repoKey, imageName, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.TagMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string, string, string) error); ok {
		r1 = rf(ctx, parentID, repoKey, imageName, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasTagsAfterName provides a mock function with given fields: ctx, repoID, filters
func (_m *TagRepository) HasTagsAfterName(ctx context.Context, repoID int64, filters types.FilterParams) (bool, error) {
	ret := _m.Called(ctx, repoID, filters)

	if len(ret) == 0 {
		panic("no return value specified for HasTagsAfterName")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, types.FilterParams) (bool, error)); ok {
		return rf(ctx, repoID, filters)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, types.FilterParams) bool); ok {
		r0 = rf(ctx, repoID, filters)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, types.FilterParams) error); ok {
		r1 = rf(ctx, repoID, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LockTagByNameForUpdate provides a mock function with given fields: ctx, repoID, name
func (_m *TagRepository) LockTagByNameForUpdate(ctx context.Context, repoID int64, name string) (bool, error) {
	ret := _m.Called(ctx, repoID, name)

	if len(ret) == 0 {
		panic("no return value specified for LockTagByNameForUpdate")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) (bool, error)); ok {
		return rf(ctx, repoID, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) bool); ok {
		r0 = rf(ctx, repoID, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, repoID, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagsPaginated provides a mock function with given fields: ctx, repoID, image, filters
func (_m *TagRepository) TagsPaginated(ctx context.Context, repoID int64, image string, filters types.FilterParams) ([]*types.Tag, error) {
	ret := _m.Called(ctx, repoID, image, filters)

	if len(ret) == 0 {
		panic("no return value specified for TagsPaginated")
	}

	var r0 []*types.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, types.FilterParams) ([]*types.Tag, error)); ok {
		return rf(ctx, repoID, image, filters)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, types.FilterParams) []*types.Tag); ok {
		r0 = rf(ctx, repoID, image, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Tag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string, types.FilterParams) error); ok {
		r1 = rf(ctx, repoID, image, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTagRepository creates a new instance of TagRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTagRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TagRepository {
	mock := &TagRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
