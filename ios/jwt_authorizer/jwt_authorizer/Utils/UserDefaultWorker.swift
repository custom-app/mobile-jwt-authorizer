//
//  UserDefaultWorker.swift
//  jwt_authorizer
//
//  Created by Lev Baklanov on 27.10.2022.
//

import Foundation

class UserDefaultsWorker {
    
    static let shared = UserDefaultsWorker()

    private static let KEY_ACCESS_TOKEN = "auth_token"
    private static let KEY_ACCESS_TOKEN_EXPIRE = "auth_token_expire"
    private static let KEY_REFRESH_TOKEN = "refresh_token"
    private static let KEY_REFRESH_TOKEN_EXPIRE = "refresh_token_expire"
    
    func saveAuthTokens(tokens: TokensInfo) {
        let defaults = UserDefaults.standard
        defaults.set(tokens.accessToken, forKey: UserDefaultsWorker.KEY_ACCESS_TOKEN)
        defaults.set(tokens.accessTokenExpire, forKey: UserDefaultsWorker.KEY_ACCESS_TOKEN_EXPIRE)
        defaults.set(tokens.refreshToken, forKey: UserDefaultsWorker.KEY_REFRESH_TOKEN)
        defaults.set(tokens.refreshTokenExpire, forKey: UserDefaultsWorker.KEY_REFRESH_TOKEN_EXPIRE)
    }
    
    func getAccessToken() -> TokenInfo {
        let defaults = UserDefaults.standard
        let token = defaults.object(forKey: UserDefaultsWorker.KEY_ACCESS_TOKEN) as? String ?? ""
        let expiresAt = defaults.object(forKey: UserDefaultsWorker.KEY_ACCESS_TOKEN_EXPIRE) as? Int64 ?? 0
        return TokenInfo(token: token, expiresAt: expiresAt)
    }
    
    func getRefreshToken() -> TokenInfo {
        let defaults = UserDefaults.standard
        let token = defaults.object(forKey: UserDefaultsWorker.KEY_REFRESH_TOKEN) as? String ?? ""
        let expiresAt = defaults.object(forKey: UserDefaultsWorker.KEY_REFRESH_TOKEN_EXPIRE) as? Int64 ?? 0
        return TokenInfo(token: token, expiresAt: expiresAt)
    }
    
    func haveAuthTokens() -> Bool {
        return !getAccessToken().token.isEmpty && !getRefreshToken().token.isEmpty
    }
    
    func dropTokens() {
        let defaults = UserDefaults.standard
        defaults.set("", forKey: UserDefaultsWorker.KEY_ACCESS_TOKEN)
        defaults.set(0 as Int64, forKey: UserDefaultsWorker.KEY_ACCESS_TOKEN_EXPIRE)
        defaults.set("", forKey: UserDefaultsWorker.KEY_REFRESH_TOKEN)
        defaults.set(0 as Int64, forKey: UserDefaultsWorker.KEY_REFRESH_TOKEN_EXPIRE)
    }
}
