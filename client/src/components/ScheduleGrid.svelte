<script>
  import { onMount } from "svelte";
  import nbaTeams from "../types/nbaTeams.js";
  import { fetchGameDates } from "../api/game_dates.js";
  import { fetchWeeklySchedule } from "../api/schedule.js";
  import { buildScheduleGrid } from "../utils/build_grid.js";
  import { formatDateHeaders } from "../utils/format_date_headers.js";
  import LoadingSpinner from "./LoadingSpinner.svelte";
  import ErrorMessage from "./ErrorMessage.svelte";

  let scheduleData = $state(null);
  let gameDates = $state(null);
  let gridData = $state(null);
  let dateHeaders = $state([]);
  let loading = $state(true);
  let error = $state(null);

  onMount(async () => {
    try {
      loading = true;
      error = null;

      const [datesResult, scheduleResult] = await Promise.all([
        fetchGameDates(),
        fetchWeeklySchedule(),
      ]);

      gameDates = datesResult;
      dateHeaders = formatDateHeaders(gameDates);
      scheduleData = scheduleResult;
      gridData = buildScheduleGrid(nbaTeams, gameDates, scheduleData);
    } catch (err) {
      console.error("Error loading schedule data:", err);
      error = err.message;
    } finally {
      loading = false;
    }
  });
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
              {#each dateHeaders as header}
                <td class="game-cell">
                  <span class="team-abbr"
                    >{gridData[team.abbr]?.[header.date] || "-"}</span
                  >
                </td>
              {/each}
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  </div>
{/if}
<!-- 
<style>
  .schedule-container {
    padding: 1rem;
    max-width: 100%;
    overflow-x: auto;
  }

  .table-wrapper {
    overflow-x: auto;
    border-radius: 8px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  }

  .schedule-grid {
    width: 100%;
    border-collapse: collapse;
    background: white;
    font-size: 0.875rem;
  }

  .schedule-grid th,
  .schedule-grid td {
    padding: 0.5rem;
    border: 1px solid #e2e8f0;
    text-align: center;
  }

  .team-header,
  .date-header {
    background-color: #f8fafc;
    font-weight: 600;
    color: #374151;
    position: sticky;
    top: 0;
    z-index: 10;
  }

  .team-header {
    left: 0;
    z-index: 11;
    min-width: 180px;
  }

  .team-name {
    background-color: #f8fafc;
    position: sticky;
    left: 0;
    z-index: 5;
    text-align: left;
    min-width: 180px;
  }

  .team-abbr {
    font-weight: 600;
    color: #1f2937;
    margin-right: 0.5rem;
  }

  .team-full-name {
    color: #6b7280;
    font-size: 0.75rem;
  }

  .game-cell {
    font-weight: 500;
    min-width: 80px;
  }

  .game-cell:empty::after {
    content: "—";
    color: #d1d5db;
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
</style> -->
