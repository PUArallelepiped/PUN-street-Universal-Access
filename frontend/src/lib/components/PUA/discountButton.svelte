<script lang="ts">
	export let text: string = 'null';
</script>

<div id="checkbox">
	<label
		><input type="checkbox" name="variety" value="" />
		<span class="round button">{text}</span></label
	>
</div>

<style>
	#checkbox input[type='checkbox'] {
		display: none;
	}
	#checkbox input:checked + .button {
		background: #487444;
		color: white;
	}
	#checkbox .button {
		display: inline-block;
		margin: 0 5px 10px 0;
		padding: 5px 10px;
		background: transparent;
		cursor: pointer;
		color: #287821;
		font-weight: bold;
		padding-top: 0px;
		padding-bottom: 0px;
	}
	#checkbox .button:hover {
		border-color: #287821;
	}
	#checkbox .round {
		border-radius: 10px;
		border: 2px solid #2d9a25;
		border-color: #287821;
	}
</style>
