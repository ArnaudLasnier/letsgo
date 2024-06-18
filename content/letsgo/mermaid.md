<!-- ## Data Model -->

<!-- ```mermaid
erDiagram
    TOURNAMENT {
        id uuid
        title text
        status enum
        started_at timestamptz
        ended_at timestamptz
    }
    PLAYER {
        id uuid
        username text
    }
    MATCH {
        id uuid
        tournament_id uuid
        parent_match_1_id uuid
        parent_match_2_id uuid
        due_at timestamptz
        opponent_1_id uuid
        opponent_1_score integer
        opponent_2_id uuid
        opponent_2_score integer
    }
    TOURNAMENT_PARTICIPATION {
        tournament_id uuid
        participant_id uuid
    }
    PLAYER ||--o{ TOURNAMENT_PARTICIPATION: participates
    TOURNAMENT_PARTICIPATION }o--|| TOURNAMENT: has_participants
    MATCH }o--|| TOURNAMENT: belongs_to
    MATCH }o--o| MATCH: originates
    MATCH }o--|{ PLAYER: plays_in
``` -->

<!-- .image data_model.svg 500 _ -->