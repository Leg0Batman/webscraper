import Foundation
import SwiftUI

class APIService: ObservableObject {
    @Published var apiKey: String = ""

    func fetchAPIKey(from url: String) {
        guard let requestURL = URL(string: "http://localhost:8080/scrape") else { return }

        var request = URLRequest(url: requestURL)
        request.httpMethod = "POST"
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")

        let body: [String: String] = ["url": url]
        request.httpBody = try? JSONSerialization.data(withJSONObject: body)

        URLSession.shared.dataTask(with: request) { data, response, error in
            if let data = data {
                do {
                    let response = try JSONDecoder().decode(Response.self, from: data)
                    DispatchQueue.main.async {
                        self.apiKey = response.result
                    }
                } catch {
                    print("Error decoding JSON: \(error)")
                }
            }
        }.resume()
    }
}

struct Response: Codable {
    let result: String
}
