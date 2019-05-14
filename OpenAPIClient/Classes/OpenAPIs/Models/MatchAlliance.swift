//
// MatchAlliance.swift
//
// Generated by openapi-generator
// https://openapi-generator.tech
//

import Foundation



public struct MatchAlliance: Codable {

    /** Score for this alliance. Will be null or -1 for an unplayed match. */
    public var score: Int
    public var teamKeys: [String]
    /** TBA team keys (eg &#x60;frc254&#x60;) of any teams playing as a surrogate. */
    public var surrogateTeamKeys: [String]?
    /** TBA team keys (eg &#x60;frc254&#x60;) of any disqualified teams. */
    public var dqTeamKeys: [String]?

    public init(score: Int, teamKeys: [String], surrogateTeamKeys: [String]?, dqTeamKeys: [String]?) {
        self.score = score
        self.teamKeys = teamKeys
        self.surrogateTeamKeys = surrogateTeamKeys
        self.dqTeamKeys = dqTeamKeys
    }

    public enum CodingKeys: String, CodingKey { 
        case score
        case teamKeys = "team_keys"
        case surrogateTeamKeys = "surrogate_team_keys"
        case dqTeamKeys = "dq_team_keys"
    }


}
