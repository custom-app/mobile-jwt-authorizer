//
//  Errors.swift
//  jwt_authorizer
//
//  Created by Lev Baklanov on 27.10.2022.
//

import Foundation

class Errors {
    
    //internal
    static let ERR_SERIALIZING_REQUEST = "error_serializing_request"
    static let ERR_CONVERTING_TO_HTTP_RESPONSE = "error_converting_response_to_http_response"
    static let ERR_PARSE_RESPONSE = "error_parsing_response"
    static let ERR_NIL_BODY = "error_nil_body"
    static let ERR_PARSE_ERROR_RESPONSE = "error_parsing_error_response"
    
    //server
    static let ERR_USER_EXIST = "user already exist"
    static let ERR_USER_NOT_EXIST = "user not exist"
    static let ERR_WRONG_CREDENTIALS = "wrong credentials"
    static let ERR_MISSING_AUTH_HEADER = "missing auth header or wrong header format"
    static let ERR_INVALID_ACCESS_TOKEN = "invalid access token"
    static let ERR_ACCESS_TOKEN_EXPIRED = "access token expired"
    static let ERR_INVALID_REFRESH_TOKEN = "invalid refresh token"
    static let ERR_REFRESH_TOKEN_EXPIRED = "refresh token expired"
    
    static func messageFor(err: String) -> String {
        switch err {
        case ERR_USER_EXIST:
            return "The user with given login already exists"
        case ERR_USER_NOT_EXIST:
            return "The user with given login doesn't exist"
        case ERR_WRONG_CREDENTIALS:
            return "Entered wrong login or password"
        default:
            return "An error has occured. Please check your internet connection and try again"
        }
    }
    
    static func isAuthError(err: String) -> Bool {
        return err == ERR_MISSING_AUTH_HEADER || err == ERR_INVALID_ACCESS_TOKEN ||
        err == ERR_INVALID_REFRESH_TOKEN || err == ERR_ACCESS_TOKEN_EXPIRED || err == ERR_REFRESH_TOKEN_EXPIRED
    }
}
