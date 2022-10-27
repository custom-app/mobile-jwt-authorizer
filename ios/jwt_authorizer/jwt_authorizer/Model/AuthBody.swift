//
//  AuthBody.swift
//  jwt_authorizer
//
//  Created by Lev Baklanov on 27.10.2022.
//

import Foundation

struct AuthBody: Codable {
    let login: String
    let password: String
}
