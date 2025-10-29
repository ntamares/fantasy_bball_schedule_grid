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

<ScheduleGrid
  {gridData}
  {dateHeaders}
  {gameDates}
  {teamGameCount}
  {loading}
  {error}
/>
<FantasyTeamRosterSelection />

<style></style>
