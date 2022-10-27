//
//  IdentifiableAlert.swift
//  jwt_authorizer
//
//  Created by Lev Baklanov on 28.10.2022.
//

import Foundation
import SwiftUI

struct IdentifiableAlert: Identifiable {
    var id: String
    var alert: () -> Alert
    
    static func build(id: String, title: String, message: String) -> IdentifiableAlert {
        return IdentifiableAlert(id: id, alert: {
            Alert(
                title: Text(LocalizedStringKey(title)),
                message: Text(message)
            )
        })
    }
    
    static func buildForError(id: String, message: String) -> IdentifiableAlert {
        return IdentifiableAlert(id: id, alert: {
            Alert(
                title: Text("An error has occured"),
                message: Text(message)
            )
        })
    }
    
    static func networkError() -> IdentifiableAlert {
        return buildForError(id: "network_err", message: "Please check your internet connection and try again")
    }
}
