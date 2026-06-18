package main

func SeedMovies() (Movies, int) {
	var m Movies
	data := []Movie{
		{"The_Shawshank_Redemption", 1994, "Drama", "Two_imprisoned_men_bond_over_a_number_of_years,_finding_solace_and_eventual_redemption_through_acts_of_common_decency", 9.3, "Film_yang_sangat_menginspirasi_dan_penuh_harapan"},
		{"The_Godfather", 1972, "Crime", "The_aging_patriarch_of_an_organized_crime_dynasty_transfers_control_of_his_clandestine_empire_to_his_reluctant_son", 9.2, "Mahakarya_sinema_yang_tak_lekang_oleh_waktu"},
		{"Inception", 2010, "Sci-Fi", "A_thief_who_steals_corporate_secrets_through_dream-sharing_technology_is_given_the_inverse_task_of_planting_an_idea_into_the_mind_of_a_CEO", 8.8, "Konsep_yang_sangat_kreatif_dan_visual_yang_memukau"},
		{"Spirited_Away", 2001, "Animation", "During_her_family_move_to_the_suburbs,_a_sullen_10-year-old_girl_wanders_into_a_world_ruled_by_gods,_witches,_and_spirits", 8.6, "Animasi_yang_magis_dan_penuh_imajinasi"},
		{"Parasite", 2019, "Thriller", "Greed_and_class_discrimination_threaten_the_newly_formed_symbiotic_relationship_between_the_wealthy_Park_family_and_the_destitute_Kim_clan", 8.6, "Satire_sosial_yang_tajam_dan_mencengangkan"},
		{"The_Dark_Knight", 2008, "Action", "When_the_menace_known_as_the_Joker_wreaks_havoc_and_chaos_on_the_people_of_Gotham,_Batman_must_accept_one_of_the_greatest_psychological_tests_of_his_ability_to_fight_injustice", 9.0, "Film_superhero_terbaik_sepanjang_masa_dengan_akting_luar_biasa"},
		{"Interstellar", 2014, "Sci-Fi", "A_team_of_explorers_travel_through_a_wormhole_in_space_in_an_attempt_to_ensure_humanity_survival.", 8.7, "Perjalanan_luar_angkasa_yang_epik_dan_penuh_emosi"},
		{"Forrest_Gump", 1994, "Drama", "The_presidencies_of_Kennedy_and_Johnson,_the_Vietnam_War,_and_other_historical_events_unfold_through_the_perspective_of_an_Alabama_man_with_a_low_IQ", 8.8, "Kisah_hidup_yang_mengharukan_dan_penuh_pelajaran"},
		{"The_Lion_King", 1994, "Animation", "Lion_cub_Simba_idolizes_his_father_King_Mufasa_and_takes_to_heart_his_teachings,_but_is_traumatized_when_his_uncle_Scar_frames_him_for_the_kings_murder", 8.5, "Animasi_klasik_Disney_yang_memukau_dan_menyentuh_hati"},
		{"Avengers_Endgame", 2019, "Action", "After_the_devastating_events_of_Infinity_War,_the_Avengers_assemble_once_more_to_reverse_Thanos_actions_and_restore_balance_to_the_universe.", 8.4, "Penutup_epik_dari_saga_Marvel_yang_memuaskan_para_penonton"},
	}
	copy(m[:], data)
	return m, len(data)
}

// code di cineReview.go
// tabMovies, n = SeedMovies()
