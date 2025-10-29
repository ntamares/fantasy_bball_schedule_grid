<script lang="js">
  import ScheduleGrid from "./ScheduleGrid.svelte";
  import nbaTeams from "../types/nbaTeams.js";
  import { onMount } from "svelte";
  import { fetchGameDates } from "../api/game_dates.js";
  import { fetchWeeklySchedule } from "../api/schedule.js";
  import { buildScheduleGrid } from "../utils/build_grid.js";
  import { formatDateHeaders } from "../utils/format_date_headers.js";
  import { getTeamGameCount } from "../utils/team_game_count.js";
  import FantasyTeamRosterSelection from "./FantasyTeamRosterSelection.svelte";
  import FreeAgentList from "./FreeAgentList.svelte";

  let scheduleData = $state(null);
  let gameDates = $state(null);
  let gridData = $state(null);
  let teamGameCount = $state(null);
  let loading = $state(true);
  let error = $state(null);
  let dateHeaders = $state([]);

  onMount(async () => {
    try {
      loading = true;
      error = null;

      const [datesResult, scheduleResult] = await Promise.all([
        fetchGameDates(),
        fetchWeeklySchedule(),
      ]);

      gameDates = datesResult;
      scheduleData = scheduleResult;
      dateHeaders = formatDateHeaders(gameDates);
      gridData = buildScheduleGrid(nbaTeams, gameDates, scheduleData);
      teamGameCount = getTeamGameCount(scheduleData);
    } catch (err) {
      console.error("Error loading schedule data:", err);
      error = err.message;
    } finally {
      loading = false;
    }
  });
</script>

<div class="container">
  <div class="schedule-section">
    <ScheduleGrid
      {gridData}
      {dateHeaders}
      {gameDates}
      {teamGameCount}
      {loading}
      {error}
    />
  </div>
  <div class="roster-section">
    <FantasyTeamRosterSelection />
  </div>
  <div class="fa-section">
    <FreeAgentList />
  </div>
</div>

<style>
  :global(body) {
    overflow: hidden;
    margin: 0;
    padding: 0;
  }

  .container {
    display: flex;
    gap: 20px;
    height: 100vh;
    width: 100vw;
    text-align: left;
    overflow: hidden;
  }

  .schedule-section {
    flex: 3;
    overflow: auto;
    max-height: 100vh;
  }

  .roster-section {
    flex: 1;
    border-left: 1px solid #ccc;
    border-right: 1px solid #ccc;
    padding: 20px;
    overflow: auto;
    max-height: 100vh;
  }

  .fa-section {
    flex: 3;
    padding: 20px;
    overflow: auto;
    max-height: 100vh;
  }
</style>
