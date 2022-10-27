//
//  ContentView.swift
//  jwt_authorizer
//
//  Created by Lev Baklanov on 27.10.2022.
//

import SwiftUI

struct ContentView: View {
    
    @StateObject
    var mainVm = MainViewModel()
    
    var body: some View {
        NavigationView {
            if mainVm.showAuthContainer {
                LoginScreen()
            } else {
                DevsScreen()
            }
        }
        .background(Color.white.ignoresSafeArea())
        .navigationViewStyle(StackNavigationViewStyle())
        .environmentObject(mainVm)
        .navigationBarTitle("", displayMode: .inline)
        .preferredColorScheme(.light)
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
