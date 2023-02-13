Rails.application.routes.draw do
  resources :quotes
  resources :users

  root "welcome#index"
end
