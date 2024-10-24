import SwiftUI

struct ContentView: View {
    @State private var url: String = ""
    @StateObject var apiService = APIService()

    var body: some View {
        VStack {
            TextField("Enter URL", text: $url)
                .textFieldStyle(RoundedBorderTextFieldStyle())
                .padding()

            Button(action: {
                apiService.fetchAPIKey(from: url)
            }) {
                Text("Scrape API Key")
            }
            .padding()

            if !apiService.apiKey.isEmpty {
                Text("API Key: \(apiService.apiKey)")
                    .padding()
            } else {
                Text("No API Key found")
                    .padding()
            }
        }
        .padding()
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}