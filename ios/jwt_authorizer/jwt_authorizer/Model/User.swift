//
//  User.swift
//  jwt_authorizer
//
//  Created by Lev Baklanov on 27.10.2022.
//

import Foundation

struct User: Codable {
    let id: Int64
    let accessToken: String
    let accessTokenExpire: Int64
    let refreshToken: String
    let refreshTokenExpire: Int64
    
    enum CodingKeys: String, CodingKey {
        case id
        case accessToken = "access_token"
        case accessTokenExpire = "access_token_expire"
        case refreshToken = "refresh_token"
        case refreshTokenExpire = "refresh_token_expire"
    }
    
    func getTokensInfo() -> TokensInfo {
        return TokensInfo(accessToken: accessToken,
                          accessTokenExpire: accessTokenExpire,
                          refreshToken: refreshToken,
                          refreshTokenExpire: refreshTokenExpire)
    }
}
