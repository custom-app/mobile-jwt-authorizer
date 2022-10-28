//
//  DevsScreen.swift
//  jwt_authorizer
//
//  Created by Lev Baklanov on 27.10.2022.
//

import SwiftUI

struct DevsScreen: View {
    
    @EnvironmentObject
    var mainVm: MainViewModel
    
    var body: some View {
        VStack(spacing: 0) {
            switch mainVm.devsProgress {
            case .finished:
                ZStack {
                    VStack(spacing: 0) {
                        ScrollView {
                            LazyVStack(spacing: 16) {
                                ForEach(mainVm.developers, id: \.id) { dev in
                                    DeveloperCard(developer: dev)
                                }
                            }
                            .padding(.vertical, 14)
                        }
                    }
                    VStack {
                        Spacer()
                        Button {
                            mainVm.getDevelopers()
                        } label: {
                            Image(systemName: "arrow.clockwise")
                                .renderingMode(.template)
                                .resizable()
                                .scaledToFit()
                                .foregroundColor(Color.white)
                                .frame(width: 30, height: 30)
                                .padding(10)
                                .background(Circle().fill(Color.black))
                                .padding(.bottom, 20)
                        }
                        .opacity(0.9)
                    }
                }
            case .error:
                Spacer()
                Text("An error occurred while loading data")
                    .padding(.horizontal, 24)
                Button {
                    mainVm.getDevelopers()
                } label: {
                    Text("Retry")
                        .foregroundColor(Color.white)
                        .fontWeight(.bold)
                        .padding(.vertical, 12)
                        .padding(.horizontal, 50)
                        .background(Color.black)
                }
                .cornerRadius(10)
                .padding(.top, 16)
                .padding(.horizontal, 24)
                Spacer()
            case .notStarted, .loading:
                Spacer()
                ProgressView()
                    .progressViewStyle(CircularProgressViewStyle(tint: .black))
                    .scaleEffect(1.6)
                    .padding(.top, 26)
                Spacer()
            }
        }
        .navigationBarTitle("Developers", displayMode: .inline)
        .toolbar {
            Button("Logout") {
                mainVm.logout()
            }
        }
        .background(Color.white.ignoresSafeArea())
        .alert(item: $mainVm.alert) { alert in
            alert.alert()
        }
        .onAppear {
            if mainVm.developers.isEmpty {
                mainVm.getDevelopers()
            }
        }
    }
}

struct DeveloperCard: View {
    
    var developer: Developer
    
    var body: some View {
        HStack(spacing: 0) {
            Image("logo_customapp")
                .resizable()
                .scaledToFit()
                .frame(width: 80, height: 40)
                .padding(10)
                .background(Color.black)
                .cornerRadius(8)
            Spacer()
            VStack(spacing: 14) {
                Text(developer.name)
                    .fontWeight(.bold)
                Text(developer.department)
                    .fontWeight(.bold)
                    .foregroundColor(departmentColor)
            }
            Spacer()
        }
        .padding(.vertical, 8)
        .padding(.horizontal, 12)
        .background(Color.white)
        .cornerRadius(8)
        .shadow(radius: 4)
        .padding(.horizontal, 20)
    }
    
    var departmentColor: Color {
        switch developer.department {
        case "Backend":
            return .teal
        case "Frontend":
            return .orange
        case "Mobile":
            return .indigo
        case "Design":
            return .red
        default:
            return .black
        }
    }
}

struct DevsScreen_Previews: PreviewProvider {
    static var previews: some View {
        DevsScreen()
    }
}
