// The Swift Programming Language
// https://docs.swift.org/swift-book

import Foundation

struct Point: Hashable, CustomStringConvertible {
    var x: Int
    var y: Int

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
            solvePart1(grid: grid)
        } catch {
            print("error \(error)")
        }
    }

    static func solvePart1(grid: [[Character]]) {
        // Used for optional unwrapping and early exit. It's the opposite of if let
        // guard condition else {} ==> if !condition {}
        guard let startX = grid[0].firstIndex(of: "S") else {
            print("No start position found")
            return
        }
        var particles: Set<Point> = [Point(x: startX, y: 0)]
        var splitCount = 0
        let maxY = grid.count - 1
        while !particles.allSatisfy({ $0.y == maxY }) {
            var newParticles: Set<Point> = []
            for particle in particles {
                guard particle.y < maxY else {
                    newParticles.insert(particle)
                    print("Particle \(particle) already at the end")
                    continue
                }
                let nextY = particle.y + 1
                if grid[nextY][particle.x] == "^" {
                    newParticles.insert(Point(x: particle.x - 1, y: nextY))
                    newParticles.insert(Point(x: particle.x + 1, y: nextY))
                    splitCount += 1
                } else {
                    newParticles.insert(Point(x: particle.x, y: nextY))
                }
            }
            particles = newParticles
        }
        print("Problem 1: \(splitCount)")
    }
}
