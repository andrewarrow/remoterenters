class WelcomeController < ApplicationController

  def index
    @quotes = Quote.all
  end

end
