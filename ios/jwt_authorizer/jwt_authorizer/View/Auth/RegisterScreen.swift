//
//  RegisterScreen.swift
//  jwt_authorizer
//
//  Created by Lev Baklanov on 27.10.2022.
//

import SwiftUI

struct RegisterScreen: View {
    
    @EnvironmentObject
    var mainVm: MainViewModel
    
    @State
    var login = ""
    @State
    var password = ""
    
    
    var body: some View {
        VStack(spacing: 0) {
            Image("logo_customapp_fill")
                .resizable()
                .scaledToFit()
                .frame(width: 240, height: 135)
                .padding(.top, 60)
            VStack(spacing: 0) {
                Text("Sign up")
                    .foregroundColor(.black)
                    .fontWeight(.bold)
                HStack {
                    Text("Login")
                        .foregroundColor(.black)
                    Spacer()
                }
                .padding(.top, 12)
                
                TextField("", text: $login.animation())
                    .foregroundColor(.black)
                    .disabled(mainVm.registerPending)
                    .padding(.vertical, 8)
                    .padding(.horizontal, 10)
                    .cornerRadius(10)
                    .overlay(RoundedRectangle(cornerRadius: 10)
                                .stroke(.black, lineWidth: 1)
                                .opacity(0.5))
                    .padding(.top, 4)
                    
                
                HStack {
                    Text("Password")
                        .foregroundColor(.black)
                    Spacer()
                }
                .padding(.top, 12)
                SecureField("", text: $password)
                    .foregroundColor(.black)
                    .font(Font.custom("montserrat-medium", size: 17))
                    .disabled(mainVm.registerPending)
                    .padding(.vertical, 8)
                    .padding(.horizontal, 10)
                    .cornerRadius(10)
                    .overlay(RoundedRectangle(cornerRadius: 10)
                                .stroke(.black, lineWidth: 1)
                                .opacity(0.5))
                    .padding(.top, 4)
                
                if mainVm.loginPending {
                    ProgressView()
                        .progressViewStyle(CircularProgressViewStyle(tint: .black))
                        .scaleEffect(1.4)
                        .padding(.top, 26)
                } else {
                    Button {
                        hideKeyboard()
                        if login.isEmpty {
                            mainVm.alert = IdentifiableAlert.build(id: "empty_login",
                                                                   title: "Invalid login",
                                                                   message: "Login can't be empty")
                            return
                        }
                        if password.count < 4 {
                            mainVm.alert = IdentifiableAlert.build(id: "empty_login",
                                                                   title: "Invalid password",
                                                                   message: "Password lenght must be 4 or more characters")
                            return
                        }
                        mainVm.register(login: login, password: password)
                    } label: {
                        Text("Register")
                            .foregroundColor(Color.white)
                            .fontWeight(.bold)
                            .padding(.vertical, 12)
                            .frame(maxWidth: .infinity)
                            .background(Color.black)
                    }
                    .disabled(mainVm.registerPending)
                    .cornerRadius(10)
                    .padding(.top, 26)
                    .padding(.horizontal, 24)
                }
            }
            .padding(.horizontal, 20)
            .padding(.top, 8)
            .padding(.bottom, 16)
            .background(Color.white)
            .cornerRadius(8)
            .shadow(radius: 4)
            .padding(.horizontal, 40)
            .padding(.top, 80)
            
            Spacer()
        }
        .navigationBarTitle("Registration", displayMode: .inline)
        .background(Color.white.ignoresSafeArea())
        .alert(item: $mainVm.alert) { alert in
            alert.alert()
        }
    }
}

struct RegisterScreen_Previews: PreviewProvider {
    static var previews: some View {
        RegisterScreen()
    }
}
