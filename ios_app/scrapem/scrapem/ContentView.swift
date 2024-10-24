import SwiftUI

struct ContentView: View {
    @State private var url: String = ""
    @State private var result: String = ""

    var body: some View {
        VStack {
            TextField("Enter GitHub URL", text: $url)
                .textFieldStyle(RoundedBorderTextFieldStyle())
                .padding()

            Button(action: {
                self.scrapeURL()
            }) {
                Text("Scrape")
            }
            .padding()

            Text(result)
                .padding()
        }
        .padding()
    }

    func scrapeURL() {
        guard let url = URL(string: "http://localhost:8080/scrape") else { return }
        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")

        let body: [String: String] = ["url": self.url]
        request.httpBody = try? JSONSerialization.data(withJSONObject: body)

        URLSession.shared.dataTask(with: request) { data, response, error in
            guard let data = data, error == nil else { return }
            if let result = try? JSONSerialization.jsonObject(with: data, options: []) as? [String: Any] {
                DispatchQueue.main.async {
                    self.result = result.description
                }
            }
        }.resume()
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
