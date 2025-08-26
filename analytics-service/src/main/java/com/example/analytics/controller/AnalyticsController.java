package com.example.analytics.controller;

import com.example.analytics.model.Expense;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

import java.util.*;

@RestController
@CrossOrigin(origins = "*") // allow local frontend
@RequestMapping("/summary")
public class AnalyticsController {

    private final String EXPENSE_SERVICE_URL = "http://expense-service:8080/expenses";
    private final String USER_SERVICE_URL = "http://user-service:5000/users";


    @GetMapping("/{userId}")
    public Map<String, Object> getSummary(@PathVariable int userId) {
        RestTemplate restTemplate = new RestTemplate();
        Expense[] expenses = restTemplate.getForObject(EXPENSE_SERVICE_URL + "/" + userId, Expense[].class);
	System.out.println("expenses " + Arrays.toString(expenses));

        double totalAmount = 0;
        Map<String, Double> categoryTotals = new HashMap<>();

        if (expenses != null) {
            for (Expense e : expenses) {
                totalAmount += e.getAmount();
                categoryTotals.put(e.getCategory(), categoryTotals.getOrDefault(e.getCategory(), 0.0) + e.getAmount());
            }
        }

        Map<String, Object> summary = new HashMap<>();
        summary.put("totalAmount", totalAmount);
        summary.put("categoryTotals", categoryTotals);
        summary.put("expenses", expenses);
        return summary;
    }
}
