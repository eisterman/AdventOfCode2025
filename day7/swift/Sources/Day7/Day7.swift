// The Swift Programming Language
// https://docs.swift.org/swift-book

import Foundation

class Point: Hashable, Equatable, CustomStringConvertible {
    // Class is REFERENCE-LIKE, while Struct is VALUE-LIKE
    var x: Int
    var y: Int
    init(x: Int, y: Int) {
        self.x = x
        self.y = y
    }

    func hash(into hasher: inout Hasher) {
        hasher.combine(x)
        hasher.combine(y)
    }

    static func == (lhs: Point, rhs: Point) -> Bool {
        lhs.x == rhs.x && lhs.y == rhs.y
    }

    var description: String {
        "(\(self.x),\(self.y))"
    }
}

@main
struct Day7 {
    static func main() {
        do {
            // let content = try String.init(contentsOfFile: "../test.data", encoding: .utf8)
            let content = try String.init(contentsOfFile: "../input.data", encoding: .utf8)
            let grid: [[Character]] = content.split(separator: "\n").map({ Array($0) })  // grid[y][x]
            var particles: [Point] = []
            // FInd initial particle
            let sx = grid[0].firstIndex(of: "S")!
            particles.append(Point(x: sx, y: 0))
            var res1 = 0
            while !particles.allSatisfy({ $0.y == grid.count - 1 }) {
                var new_parts: [Point] = []
                for i in 0..<particles.count {
                    let p = particles[i]
                    if grid[p.y + 1][p.x] == "^" {
                        new_parts.append(Point(x: p.x + 1, y: p.y + 1))
                        p.y += 1
                        p.x -= 1
                        res1 += 1
                    } else {
                        p.y += 1
                    }
                }
                particles.append(contentsOf: new_parts)
                // Remove duplicates
                particles = Array(Set(particles))
            }
            print("Problem 1: \(res1)")
        } catch {
            print("error \(error)")
        }
    }
}
