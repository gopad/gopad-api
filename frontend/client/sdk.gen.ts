// This file is auto-generated by @hey-api/openapi-ts

import type {
  Options as ClientOptions,
  TDataShape,
  Client,
} from '@hey-api/client-fetch'
import type {
  RequestProviderData,
  RequestProviderError,
  CallbackProviderData,
  CallbackProviderError,
  ListProvidersData,
  ListProvidersResponse,
  RedirectAuthData,
  RedirectAuthResponse,
  RedirectAuthError,
  LoginAuthData,
  LoginAuthResponse,
  LoginAuthError,
  RefreshAuthData,
  RefreshAuthResponse,
  RefreshAuthError,
  VerifyAuthData,
  VerifyAuthResponse,
  VerifyAuthError,
  TokenProfileData,
  TokenProfileResponse,
  TokenProfileError,
  ShowProfileData,
  ShowProfileResponse,
  ShowProfileError,
  UpdateProfileData,
  UpdateProfileResponse,
  UpdateProfileError,
  ListGroupsData,
  ListGroupsResponse,
  ListGroupsError,
  CreateGroupData,
  CreateGroupResponse,
  CreateGroupError,
  DeleteGroupData,
  DeleteGroupResponse,
  DeleteGroupError,
  ShowGroupData,
  ShowGroupResponse,
  ShowGroupError,
  UpdateGroupData,
  UpdateGroupResponse,
  UpdateGroupError,
  DeleteGroupFromUserData,
  DeleteGroupFromUserResponse,
  DeleteGroupFromUserError,
  ListGroupUsersData,
  ListGroupUsersResponse,
  ListGroupUsersError,
  AttachGroupToUserData,
  AttachGroupToUserResponse,
  AttachGroupToUserError,
  PermitGroupUserData,
  PermitGroupUserResponse,
  PermitGroupUserError,
  ListUsersData,
  ListUsersResponse,
  ListUsersError,
  CreateUserData,
  CreateUserResponse,
  CreateUserError,
  DeleteUserData,
  DeleteUserResponse,
  DeleteUserError,
  ShowUserData,
  ShowUserResponse,
  ShowUserError,
  UpdateUserData,
  UpdateUserResponse,
  UpdateUserError,
  DeleteUserFromGroupData,
  DeleteUserFromGroupResponse,
  DeleteUserFromGroupError,
  ListUserGroupsData,
  ListUserGroupsResponse,
  ListUserGroupsError,
  AttachUserToGroupData,
  AttachUserToGroupResponse,
  AttachUserToGroupError,
  PermitUserGroupData,
  PermitUserGroupResponse,
  PermitUserGroupError,
} from './types.gen'
import { client as _heyApiClient } from './client.gen'

export type Options<
  TData extends TDataShape = TDataShape,
  ThrowOnError extends boolean = boolean,
> = ClientOptions<TData, ThrowOnError> & {
  /**
   * You can provide a client instance returned by `createClient()` instead of
   * individual options. This might be also useful if you want to implement a
   * custom client.
   */
  client?: Client
  /**
   * You can pass arbitrary values through the `meta` object. This can be
   * used to access values that aren't defined as part of the SDK function.
   */
  meta?: Record<string, unknown>
}

/**
 * Request the redirect to defined provider
 */
export const requestProvider = <ThrowOnError extends boolean = false>(
  options: Options<RequestProviderData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).get<
    unknown,
    RequestProviderError,
    ThrowOnError
  >({
    url: '/auth/{provider}/request',
    ...options,
  })
}

/**
 * Callback to parse the defined provider
 */
export const callbackProvider = <ThrowOnError extends boolean = false>(
  options: Options<CallbackProviderData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).get<
    unknown,
    CallbackProviderError,
    ThrowOnError
  >({
    url: '/auth/{provider}/callback',
    ...options,
  })
}

/**
 * Fetch the available auth providers
 */
export const listProviders = <ThrowOnError extends boolean = false>(
  options?: Options<ListProvidersData, ThrowOnError>
) => {
  return (options?.client ?? _heyApiClient).get<
    ListProvidersResponse,
    unknown,
    ThrowOnError
  >({
    url: '/auth/providers',
    ...options,
  })
}

/**
 * Retrieve real token after redirect
 */
export const redirectAuth = <ThrowOnError extends boolean = false>(
  options: Options<RedirectAuthData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).post<
    RedirectAuthResponse,
    RedirectAuthError,
    ThrowOnError
  >({
    url: '/auth/redirect',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Authenticate an user by credentials
 */
export const loginAuth = <ThrowOnError extends boolean = false>(
  options: Options<LoginAuthData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).post<
    LoginAuthResponse,
    LoginAuthError,
    ThrowOnError
  >({
    url: '/auth/login',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Refresh an auth token before it expires
 */
export const refreshAuth = <ThrowOnError extends boolean = false>(
  options?: Options<RefreshAuthData, ThrowOnError>
) => {
  return (options?.client ?? _heyApiClient).get<
    RefreshAuthResponse,
    RefreshAuthError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/auth/refresh',
    ...options,
  })
}

/**
 * Verify validity for an authentication token
 */
export const verifyAuth = <ThrowOnError extends boolean = false>(
  options?: Options<VerifyAuthData, ThrowOnError>
) => {
  return (options?.client ?? _heyApiClient).get<
    VerifyAuthResponse,
    VerifyAuthError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/auth/verify',
    ...options,
  })
}

/**
 * Retrieve an unlimited auth token
 */
export const tokenProfile = <ThrowOnError extends boolean = false>(
  options?: Options<TokenProfileData, ThrowOnError>
) => {
  return (options?.client ?? _heyApiClient).get<
    TokenProfileResponse,
    TokenProfileError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/profile/token',
    ...options,
  })
}

/**
 * Fetch profile details of the personal account
 */
export const showProfile = <ThrowOnError extends boolean = false>(
  options?: Options<ShowProfileData, ThrowOnError>
) => {
  return (options?.client ?? _heyApiClient).get<
    ShowProfileResponse,
    ShowProfileError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/profile/self',
    ...options,
  })
}

/**
 * Update your own profile information
 */
export const updateProfile = <ThrowOnError extends boolean = false>(
  options: Options<UpdateProfileData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).put<
    UpdateProfileResponse,
    UpdateProfileError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/profile/self',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Fetch all available groups
 */
export const listGroups = <ThrowOnError extends boolean = false>(
  options?: Options<ListGroupsData, ThrowOnError>
) => {
  return (options?.client ?? _heyApiClient).get<
    ListGroupsResponse,
    ListGroupsError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/groups',
    ...options,
  })
}

/**
 * Create a new group
 */
export const createGroup = <ThrowOnError extends boolean = false>(
  options: Options<CreateGroupData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).post<
    CreateGroupResponse,
    CreateGroupError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/groups',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Delete a specific group
 */
export const deleteGroup = <ThrowOnError extends boolean = false>(
  options: Options<DeleteGroupData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).delete<
    DeleteGroupResponse,
    DeleteGroupError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/groups/{group_id}',
    ...options,
  })
}

/**
 * Fetch a specific group
 */
export const showGroup = <ThrowOnError extends boolean = false>(
  options: Options<ShowGroupData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).get<
    ShowGroupResponse,
    ShowGroupError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/groups/{group_id}',
    ...options,
  })
}

/**
 * Update a specific group
 */
export const updateGroup = <ThrowOnError extends boolean = false>(
  options: Options<UpdateGroupData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).put<
    UpdateGroupResponse,
    UpdateGroupError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/groups/{group_id}',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Unlink a user from group
 */
export const deleteGroupFromUser = <ThrowOnError extends boolean = false>(
  options: Options<DeleteGroupFromUserData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).delete<
    DeleteGroupFromUserResponse,
    DeleteGroupFromUserError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/groups/{group_id}/users',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Fetch all users attached to group
 */
export const listGroupUsers = <ThrowOnError extends boolean = false>(
  options: Options<ListGroupUsersData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).get<
    ListGroupUsersResponse,
    ListGroupUsersError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/groups/{group_id}/users',
    ...options,
  })
}

/**
 * Attach a user to group
 */
export const attachGroupToUser = <ThrowOnError extends boolean = false>(
  options: Options<AttachGroupToUserData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).post<
    AttachGroupToUserResponse,
    AttachGroupToUserError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/groups/{group_id}/users',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Update user perms for group
 */
export const permitGroupUser = <ThrowOnError extends boolean = false>(
  options: Options<PermitGroupUserData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).put<
    PermitGroupUserResponse,
    PermitGroupUserError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/groups/{group_id}/users',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Fetch all available users
 */
export const listUsers = <ThrowOnError extends boolean = false>(
  options?: Options<ListUsersData, ThrowOnError>
) => {
  return (options?.client ?? _heyApiClient).get<
    ListUsersResponse,
    ListUsersError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/users',
    ...options,
  })
}

/**
 * Create a new user
 */
export const createUser = <ThrowOnError extends boolean = false>(
  options: Options<CreateUserData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).post<
    CreateUserResponse,
    CreateUserError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/users',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Delete a specific user
 */
export const deleteUser = <ThrowOnError extends boolean = false>(
  options: Options<DeleteUserData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).delete<
    DeleteUserResponse,
    DeleteUserError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/users/{user_id}',
    ...options,
  })
}

/**
 * Fetch a specific user
 */
export const showUser = <ThrowOnError extends boolean = false>(
  options: Options<ShowUserData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).get<
    ShowUserResponse,
    ShowUserError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/users/{user_id}',
    ...options,
  })
}

/**
 * Update a specific user
 */
export const updateUser = <ThrowOnError extends boolean = false>(
  options: Options<UpdateUserData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).put<
    UpdateUserResponse,
    UpdateUserError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/users/{user_id}',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Unlink a group from user
 */
export const deleteUserFromGroup = <ThrowOnError extends boolean = false>(
  options: Options<DeleteUserFromGroupData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).delete<
    DeleteUserFromGroupResponse,
    DeleteUserFromGroupError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/users/{user_id}/groups',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Fetch all groups attached to user
 */
export const listUserGroups = <ThrowOnError extends boolean = false>(
  options: Options<ListUserGroupsData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).get<
    ListUserGroupsResponse,
    ListUserGroupsError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/users/{user_id}/groups',
    ...options,
  })
}

/**
 * Attach a group to user
 */
export const attachUserToGroup = <ThrowOnError extends boolean = false>(
  options: Options<AttachUserToGroupData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).post<
    AttachUserToGroupResponse,
    AttachUserToGroupError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/users/{user_id}/groups',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Update group perms for user
 */
export const permitUserGroup = <ThrowOnError extends boolean = false>(
  options: Options<PermitUserGroupData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).put<
    PermitUserGroupResponse,
    PermitUserGroupError,
    ThrowOnError
  >({
    security: [
      {
        name: 'X-API-Key',
        type: 'apiKey',
      },
      {
        scheme: 'bearer',
        type: 'http',
      },
      {
        scheme: 'basic',
        type: 'http',
      },
    ],
    url: '/users/{user_id}/groups',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}
