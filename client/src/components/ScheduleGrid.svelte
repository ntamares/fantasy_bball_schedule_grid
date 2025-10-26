<script>
  import nbaTeams from "../types/nbaTeams.js";
  import LoadingSpinner from "./LoadingSpinner.svelte";
  import ErrorMessage from "./ErrorMessage.svelte";

  let { gridData, dateHeaders, gameDates, teamGameCount, loading, error } =
    $props();

  function getGameCountClass(gameCount) {
    if (gameCount <= 2) return "game-count-low";
    if (gameCount === 3) return "game-count-medium";
    return "game-count-high";
  }
</script>

{#if loading}
  <LoadingSpinner />
{:else if error}
  <ErrorMessage message={error} />
{:else if gridData && dateHeaders.length > 0}
  <div class="schedule-container">
    <h2>NBA Weekly Schedule</h2>
    <div class="table-wrapper">
      <table class="schedule-grid">
        <thead>
          <tr>
            <th class="team-header">Team</th>
            <th class="game-count-header"># of Games</th>
            {#each dateHeaders as header}
              <th class="date-header">{header}</th>
            {/each}
          </tr>
        </thead>
        <tbody>
          {#each nbaTeams as team}
            <tr>
              <td class="team-name">
                <span class="team-full-name">{team.name}</span>
              </td>
              <td class="game-count">
                <span
                  class:game-count-low={teamGameCount[team.name] <= 2}
                  class:game-count-medium={teamGameCount[team.name] === 3}
                  class:game-count-high={teamGameCount[team.name] >= 4}
                >
                  {teamGameCount[team.name]}
                </span>
              </td>
              {#each gameDates as date}
                <td class="game-cell">
                  <span class="team-abbr">
                    {gridData[team.name]?.[date] || "-"}
                  </span>
                </td>
              {/each}
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  </div>
{/if}

<style>
  table,
  th,
  td {
    border: 1px solid black;
  }

  .schedule-container {
    padding: 1rem;
    max-width: 100%;
    overflow-x: auto;
  }

  .table-wrapper {
    overflow-x: auto;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
    border: 1px solid black; /* Example: 1px solid black border */
  }

  .schedule-grid {
    width: 100%;
    border-collapse: collapse;
    background: rgb(182, 201, 198);
    font-size: 0.875rem;
  }

  .schedule-grid th,
  .schedule-grid td {
    padding: 0.15rem;
    border: 1px solid #e2e8f0;
    text-align: center;
  }

  .team-header,
  .date-header,
  .game-count-header {
    background-color: #307dca;
    font-weight: 600;
    color: black;
    position: sticky;
    top: 0;
    z-index: 10;
    font-size: 1.25em;
  }

  .team-header {
    left: 0;
    z-index: 11;
    min-width: 180px;
  }

  .team-name {
    background-color: #a4b0bd;
    position: sticky;
    left: 0;
    z-index: 5;
    text-align: left;
    min-width: 180px;
  }

  .game-count {
    background-color: #bcc6d4;
    color: black;
    font-size: 1.25em;
  }

  .team-abbr {
    font-weight: 550;
    color: #1f2937;
    margin-right: 0.5rem;
    font-size: 1.25em;
  }

  .team-full-name {
    color: black;
    font-size: 1.15rem;
  }

  .game-cell {
    font-weight: 500;
    min-width: 80px;
  }

  .game-cell:empty::after {
    content: "—";
    color: #d1d5db;
  }

  .game-count-low {
    color: red;
  }
  .game-count-medium {
    color: black;
  }
  .game-count-high {
    color: green;
  }

  tbody tr:hover .team-name {
    background-color: rgba(59, 130, 246, 0.2); /* Blue overlay on #f8fafc */
  }

  tbody tr:hover .game-count {
    background-color: rgba(59, 130, 246, 0.15); /* Lighter overlay on #bcc6d4 */
  }

  tbody tr:hover .game-cell {
    background-color: rgba(59, 130, 246, 0.1); /* Even lighter overlay */
  }

  @media (max-width: 768px) {
    .schedule-grid {
      font-size: 0.75rem;
    }

    .schedule-grid th,
    .schedule-grid td {
      padding: 0.25rem;
    }

    .team-name {
      min-width: 120px;
    }

    .game-cell {
      min-width: 60px;
    }
  }
</style>
