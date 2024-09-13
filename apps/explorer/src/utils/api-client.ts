import axios from 'axios';

export const apiClient = axios.create({
	baseURL: process.env.NEXT_PUBLIC_API_URL,
});

export const carbonMonsterClient = axios.create({
	baseURL: process.env.NEXT_PUBLIC_CARBON_MONSTER_API_URL,
});
